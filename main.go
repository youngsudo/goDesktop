package main

import (
	"gogui/server"
	"os"
	"os/exec"
	"os/signal"
)

func main() {
	go server.Run()
	cmd := startBrowser()
	chSignal := listenToInterrupt()

	// 等待中断信号
	select {
	case <-chSignal:
		cmd.Process.Kill()
	}
}

func startBrowser() *exec.Cmd {
	// 找到chrome路径
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	// 创建命令
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:27149/static/index.html")
	// 启动 一个进程,启动进程比启动一个go协程慢得多
	cmd.Start()
	return cmd
}

/* 监听中断信号	*/
func listenToInterrupt() chan os.Signal {
	//   os.Signal操作系统信号, 1 是缓存
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
