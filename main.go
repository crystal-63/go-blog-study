package main

import (
	"go-blog/common"
	"go-blog/server"
)

func init() {
	common.LoadTemplate()
}

func main() {
	// 程序入口，一个项目只能有一个入口
	//	web程序， http协议：ip  port
	server.App.Start("127.0.0.1", "8077")
}
