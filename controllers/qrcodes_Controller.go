package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

// 添加二维码接口
/*
GET /api/v1/qrcodes
思路:
	1. 获取文本内容
	2. 将文本转为图片 (用库 qrcode)
	3. 将图片写入HTTP响应
	x = http://ip1;27149/upload/... .txt
	http://ip1:27149/static/downloads?url=x

	GET /api/v1/qrcodes?content=http%3A%2F%2F192.168.244.1%3A27149%2Fstatic%2Fdownloads%3Ftype%3Dtext%26url%3Dhttp%3A%2F%2F192.168.244.1%3A27149%252Fuploads%252Fdfdd8bf3-1b64-40b7-ab12-ace74752f26e.txt"
*/
func QrcodesController(c *gin.Context) {
	if content := c.Query("content"); content != "" {
		png, err := qrcode.Encode(content, qrcode.Medium, 256)
		if err != nil {
			log.Fatal(err)
		}
		c.Data(http.StatusOK, "image/png", png)
	} else {
		c.Status(http.StatusBadRequest)
	}
}
