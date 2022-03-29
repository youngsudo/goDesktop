package controllers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetUploadsDir() (uploads string) {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	if err != nil {
		log.Fatal(err)
	}
	uploads = filepath.Join(dir, "uploads")
	return
}

//
/* 文件下载
GET/uploads/:path
思路:
	1,将网络路径:path变成本地绝对路径
	2,读取本地文件,写到HTTP响应里
*/
func UploadsController(c *gin.Context) {
	if path := c.Param("path"); path != "" {
		target := filepath.Join(GetUploadsDir(), path)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary") // 内容的编码:二进制
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.Header("Content-Type", "application/octet-stream")
		c.File(target) // 给前端发送一个文件
	} else {
		c.Status(http.StatusNotFound)
	}
}
