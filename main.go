package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"huobi_Golang/api"
	"huobi_Golang/api/utils"
)

func main() {

	//gin.DefaultWriter = utils.Loger()
	//gin.SetMode(gin.DebugMode)
	//log.SetOutput(gin.DefaultWriter)
	router := gin.Default()
	//session中间件
	router.Use(sessions.Sessions("mysession", utils.Store))
	//login_token校验登陆中间件
	router.Use(utils.CheckLogin)

	app.Router(router)

	router.Run(":9800")

}
