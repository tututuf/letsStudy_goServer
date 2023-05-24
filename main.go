package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tututuf/letsStudy_goServer/router"
)

func main() {
	r := gin.Default()
	router.SetUserRouter(r)
	r.Run()
}
