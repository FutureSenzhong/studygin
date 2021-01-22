package routes

import (
	"github.com/gin-gonic/gin"
	"studygin/handler"
)


func Routers(r *gin.Engine) {
	//r.Static("/assets", "./assets")
	//r.SetHTMLTemplate(html)
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	//路由分组设置
	v1 := r.Group("v1")
	{
		v1.GET("/ping", handler.Ping)
		v1.GET("/someJSON", handler.SomeJson)
		v1.GET("/index", handler.Index)

		v1.GET("/json", handler.Json)

		v1.GET("/purejson", handler.PureJson)
		v1.POST("/post", handler.Post)
		v1.GET("/SecureJSON", handler.SecureJSON)

		v1.POST("/upload", handler.Upload)
		v1.GET("/someDataFromReader", handler.GetReader)

		v1.GET("/testing", handler.StartPage)

		v1.GET("/longAsync", handler.LongAsync)
		v1.GET("/longSync", handler.LongSync)
	}

}
