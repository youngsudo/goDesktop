package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// serveWs handles websocket requests from the peer.
// 处理websocket请求
func wshandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// 首先升级websocket协议
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 没有升级失败则初始化一个client信息,表示有一个客户正在连接
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	// 注册用户
	client.hub.register <- client

	// 允许通过在中执行所有工作来收集调用者引用的内存
	// 新的goroutines。
	go client.writePump()
	go client.readPump()
}
func HttpController(c *gin.Context, hub *Hub) {
	wshandler(hub, c.Writer, c.Request)
}
