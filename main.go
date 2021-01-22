package main

import (
	"studygin/routes"
	"studygin/server"
)


//var html = template.Must(template.New("https").Parse(`
//<html>
//<head>
//  <title>Https Test</title>
//  <script src="/assets/app.js"></script>
//</head>
//<body>
//  <h1 style="color:red;">Welcome, Ginner!</h1>
//</body>
//</html>
//`))


func main() {

	// gin服务初始化配置
	route := server.ServerIni()

	//路由函数
	routes.Routers(route)

	//启动服务
	_ = route.Run("0.0.0.0:8081") // 监听并在 0.0.0.0:8080 上启动服务
}