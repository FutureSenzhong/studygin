package server

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func EngineIni() *gin.Engine {
	//禁止日志颜色
	//gin.DisableConsoleColor()

	// 强制日志颜色化
	gin.ForceConsoleColor()

	// 自定义路由日志格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("路由日志： %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	// 记录到文件。
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	return gin.Default()
}
