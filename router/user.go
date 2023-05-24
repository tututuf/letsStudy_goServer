package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tututuf/letsStudy_goServer/controllers"
)

func SetUserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	userRouter.POST("/getInfo", controllers.GetUserInfo)
	userRouter.POST("/login", controllers.Login)
}
