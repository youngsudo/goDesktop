package main

import (
	"gogui/cfg"
	"gogui/server"
	"os"
	"os/exec"
	"os/signal"
)

func main() {
	// golang windows下 调用外部程序隐藏cmd窗口
	// go build -ldflags -H=windowsgui
	chChromeDie := make(chan struct{})
	chBackendDie := make(chan struct{})
	chSignal := listenToInterrupt()
	go server.Run()
	go startBrowser(chChromeDie, chBackendDie)

	// 等待中断信号
	for {
		select {
		case <-chSignal:
			chBackendDie <- struct{}{}
		case <-chChromeDie:
			os.Exit(0)
		}
	}
}

func startBrowser(chChromeDie, chBackendDie chan struct{}) {
	// 找到chrome路径
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	// 创建命令	实现执行外部命令以及和外部命令交互https://colobu.com/2020/12/27/go-with-os-exec/
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+cfg.GetPort()+"/static/index.html")
	// 启动 一个进程,启动进程比启动一个go协程慢得多
	cmd.Start()

	go func() {
		<-chBackendDie
		cmd.Process.Kill()
	}()
	go func() {
		cmd.Wait()
		chChromeDie <- struct{}{}
	}()
}

/* 监听中断信号	*/
func listenToInterrupt() chan os.Signal {
	//   os.Signal操作系统信号, 1 是缓存
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
