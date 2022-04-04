package main

import (
	"embed"
	"local/cfg"
	"local/server"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/zserge/lorca"
)

//go:embed frontend/dist/*
var FS embed.FS

var config *cfg.Config

func init() {
	config = cfg.GetConfig()
}

func main() {
	chBackendStarted := make(chan struct{})          // 后端程序退出信号
	go server.Run(config.Port, FS, chBackendStarted) // 启动后台服务
	<-chBackendStarted                               // 等待后端服务启动
	// 启动浏览器
	ui, err := lorca.New("", "", config.Width, config.Height)
	if err != nil {
		log.Fatal(err)
	}
	err = ui.Load("http://127.0.0.1:" + strconv.Itoa(config.Port) + "/static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close() // 在程序退出前关闭窗口

	// 等待中断信号
	chSignal := listenToInterrupt()
	select {
	case <-chSignal:
	case <-ui.Done():
	}

}

// 监听键盘ctrl+c信号
func listenToInterrupt() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt) // 注册中断信号
	return ch
}
