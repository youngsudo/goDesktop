package ws

import (
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients. 广播,触发一个事件
	clients map[*Client]bool

	// Inbound messages from the clients. 注册,监听
	broadcast chan []byte

	// Register requests from the clients. 不监听
	register chan *Client

	// Unregister requests from clients. 统计有多少人监听
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

var once sync.Once
var singleton *Hub

func (h *Hub) Run() {
	for {
		select {
		// 如果有人监听,就把它放在我的客户(client)里面
		case client := <-h.register:
			h.clients[client] = true
		// 如果有人取消注册了(监听),就把它从客户里面删除
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		// 如果有广播消息,遍历所有的客户,把消息发送给客户
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
