package main

import (
	"os/exec"
)

func main() {
	// 找到chrome路径
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	// 创建命令
	cmd := exec.Command(chromePath, "--app=https://www.baidu.com")
	// 启动
	cmd.Start()
}
