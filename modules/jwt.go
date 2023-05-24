package modules

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tututuf/letsStudy_goServer/env"
	"github.com/tututuf/letsStudy_goServer/interfaces"
)

type UserClaims struct {
	*interfaces.UserInfo
	*jwt.RegisteredClaims
}

var Secret = []byte("look-living")

func NewToken(userInfo *interfaces.UserInfo) (string, error) {
	maxAge := env.Config.MustInt("token", "maxAge")
	tokenExpireDuration := time.Hour * time.Duration(maxAge)
	numberDate := time.Now().Add(tokenExpireDuration)
	claims := &UserClaims{
		userInfo,
		&jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: numberDate,
			}, // 过期时间
			Issuer: "look-living", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString(Secret); err != nil {
		return "", err
	} else {
		fmt.Println(tokenString)
		return tokenString, nil
	}
}

func ParseToken(ctx *gin.Context) (*interfaces.UserInfo, error) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		auth = ctx.Request.Header.Get("Sec-WebSocket-Protocol")
	}
	token, err := jwt.ParseWithClaims(auth, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		log.Println("error", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		fmt.Printf("456,%-v", claims)
		return claims.UserInfo, nil
	}
	return nil, errors.New("token无法解析")
}
