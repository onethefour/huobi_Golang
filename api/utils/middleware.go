package utils

import (
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func CheckLogin(ctx *gin.Context) {
	var err error
	v, err := url.ParseQuery(ctx.Request.URL.RawQuery)
	if err != nil {
		log.Println(err.Error())
		ctx.Abort()
		ctx.String(200, err.Error())
		return
	}

	v.Set("islogin", "0")
	//是否登录
	session := sessions.Default(ctx)
	account := session.Get("account")
	if account != nil {
		v.Set("islogin", "1")
	}
	//结果写入url
	ctx.Request.URL.RawQuery = v.Encode()
	ctx.Next()
	return
}
