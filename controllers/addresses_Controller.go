package controllers

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 实现接口2: 获取局域网IP

/*
思路:
	1,获取电脑在各个局域网的IP地址
	2,转为JSON写入HTTP响应
*/
func AddressesController(c *gin.Context) {
	// InterfaceAddrs获取本地IP
	addrs, _ := net.InterfaceAddrs()
	var result []string
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		// address.类型断言,断言net是一个地址ip 是一个(*net.IPNet)类型
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"addresses": result})
}
