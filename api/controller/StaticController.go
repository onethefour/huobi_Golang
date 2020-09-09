package controller

//go-bindata-assetfs -pkg bindata ./element-html/dist/... -o ./app/bindata/bindata.go
import (
	"github.com/gin-gonic/gin"
	"huobi_Golang/api/bindata"
)

type StaticController struct{}

func (ctr *StaticController) Router(r *gin.Engine) {
	//r.GET("/", func(c *gin.Context) {
	//	c.File("./element-html/dist/index.html")
	//})
	//
	//r.GET("/:filename", func(c *gin.Context) {
	//	filename := c.Param("filename")
	//	c.File("./element-html/dist/"+filename)
	//})

	r.GET("/", func(c *gin.Context) {
		data, _ := bindata.Asset("element-html/dist/index.html")
		c.Writer.Write(data)
	})

	r.GET("/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		data, _ := bindata.Asset("element-html/dist/" + filename)
		c.Writer.Write(data)
	})

	// Write asset
	//c.Writer.Write(data)
	//r.StaticFS("/", http.Dir("./element-html/dist"))
	//r.StaticFS("/front", bindata.())
}
