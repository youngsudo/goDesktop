package main

import (
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		gin.SetMode(gin.ReleaseMode)
		router := gin.Default()

		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "<h1>Hello World</h1>")
		})
		router.Run(":8080")
	}()
	time.Sleep(time.Second * 3)

	// 找到chrome路径
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	// 创建命令
	cmd := exec.Command(chromePath, "--app=https://www.baidu.com")
	// 启动
	cmd.Start()

	select {}
}
