package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var FS embed.FS

func main() {
	go func() { // gin协程
		gin.SetMode(gin.DebugMode) //设置模式 ReleaseMode生产模式,DebugMode开发模式
		router := gin.Default()

		// router.GET("/", func(c *gin.Context) {
		// 	c.Writer.Write([]byte("abcd"))
		// })
		staticFiles, _ := fs.Sub(FS, "frontend/dist")
		router.StaticFS("/static", http.FS(staticFiles))

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
		router.Run(":8080")
	}()

	// 找到chrome路径
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	// 创建命令
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080/static/index.html")
	// 启动 一个进程,启动进程比启动一个go协程慢得多
	cmd.Start()

	/* 处理中断信号	*/
	//   os.Signal操作系统信号, 1 是缓存
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	select {
	case <-chSignal:
		cmd.Process.Kill()
	}
}
