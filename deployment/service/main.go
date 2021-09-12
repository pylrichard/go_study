package main

import (
	"go/go_study/deployment/service/builder"
)

/*
	Linux部署分为以下几种方式：
		- 使用nohup命令
		- 使用supervisord管理
		- 使用systemd管理
 */
func main() {
	//simple.Simple()
	builder.Build()
}