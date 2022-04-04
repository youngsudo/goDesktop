package server

import (
	"embed"
	"io/fs"
	c "local/controllers"
	"local/server/ws"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var hub *ws.Hub

func init() {
	hub = ws.NewHub()
	go hub.Run()
}

func Run(port int, FS embed.FS, chBackendStarted chan struct{}) {
	gin.SetMode(gin.ReleaseMode) //设置模式 ReleaseMode生产模式,DebugMode开发模式
	router := gin.Default()
	gin.DisableConsoleColor()
	router.SetTrustedProxies([]string{strconv.Itoa(port)})

	// 静态文件路由
	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	router.StaticFS("/static", http.FS(staticFiles))
	// 动态路由
	router.POST("/api/v1/texts", c.TextsController)
	router.GET("/api/v1/addresses", c.AddressesController)
	router.GET("/uploads/:path", c.UploadsController)
	router.GET("/api/v1/qrcodes", c.QrcodesController)
	router.POST("/api/v1/files", c.FilesController)

	router.GET("/ws", func(c *gin.Context) {
		ws.HttpController(c, hub)
	})
	// 没有路由时,走这最后一个路由
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path               // 获取用户访问路径
		if strings.HasPrefix(path, "/static/") { // 以static开头的,说明用户想访问的是静态文件
			reader, err := staticFiles.Open("index.html") // 打开index.html
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close() // 读完index.html后关闭文件
			stat, err := reader.Stat()
			if err != nil {
				log.Fatal(err)
			}
			c.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
		} else {
			// 不是以static开头的,说明用户想访问的是动态文件
			c.Status(http.StatusNotFound)
		}
	})
	go func() {
		chBackendStarted <- struct{}{}
	}()
	runErr := router.Run(":" + strconv.Itoa(port))
	if runErr != nil {
		log.Fatal(runErr)
	}
}
