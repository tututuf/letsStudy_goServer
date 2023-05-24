package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tututuf/letsStudy_goServer/interfaces"
)

func GetUserInfo(ctx *gin.Context) {

}

func Login(ctx *gin.Context) {
	user := &interfaces.UserInfo{
		Name:     "张三",
		Age:      18,
		Username: "zhansan",
	}
	ctx.JSON(200, user)
}
