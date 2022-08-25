package server

import (
	"github.com/go-demo-1/define"
	"github.com/go-demo-1/server/routers"
)

// Run 启动gin
func Run() {
	// 设置路由
	r := routers.SetRouters()

	// 启动服务
	r.Run(":" + define.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
