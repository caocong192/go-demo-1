package routers

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/go-demo-1/controller"
	"github.com/go-demo-1/server/ws"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

//go:embed frontend/dist/*
var FS embed.FS

// SetRouters 设置路由
func SetRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	r := gin.Default()

	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	r.StaticFS("/static", http.FS(staticFiles))

	// 支持websocket
	hub := ws.NewHub()
	go hub.Run()

	r.GET("/ws", func(c *gin.Context) {
		ws.HttpController(c, hub)
	})

	r.GET("/uploads/:path", controller.DownloadController)

	api := r.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.POST("/texts", controller.TextsController)
			v1.GET("/addresses", controller.AddressesController)
			v1.GET("/qrcodes", controller.QrcodesController)
			v1.POST("/files", controller.FilesController)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/static/") {
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			stat, err := reader.Stat()
			if err != nil {
				log.Fatal(err)
			}
			c.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	return r
}
