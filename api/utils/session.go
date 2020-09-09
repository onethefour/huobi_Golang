package utils

import (
	"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/memstore"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/cookie"
)

var Store sessions.Store

func init() {
	//Store = memstore.NewStore([]byte("secret"))
	Store = cookie.NewStore([]byte("secret"))
}

/* 使用示例
r.GET("/incr", func(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()
	c.JSON(200, gin.H{"count": count})
})

*/
