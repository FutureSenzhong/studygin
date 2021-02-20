package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"studygin/handler"
	"studygin/middleware"
)

func Routers(r *gin.Engine) {
	//r.Static("/assets", "./assets")
	//r.SetHTMLTemplate(html)
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	//配置favicon.ico
	// 静态资源加载，本例为css,js以及资源图片
	r.StaticFS("/static", http.Dir(""))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("/ping", middleware.MyLimit(), handler.Ping)
	//路由分组设置
	v1 := r.Group("v1")
	{
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

	sql := r.Group("sql")
	{
		//数据库的CRUD--->gin的 post、get、put、delete方法
		sql.POST("/insert", handler.InsertData)   //添加数据
		sql.GET("/get", handler.GetData)          //查询数据（单条记录）
		sql.GET("/mulget", handler.GetMulData)    //查询数据（多条记录）
		sql.PUT("/update", handler.UpdateData)    //更新数据
		sql.DELETE("/delete", handler.DeleteData) //删除数据
	}
}
