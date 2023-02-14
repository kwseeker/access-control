package main

import (
	"casbin-gin/cmd/config"
	"casbin-gin/cmd/server"
	"casbin-gin/common/component/casbin"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//加载配置
	config.Setup("conf/settings.yml")
	//安装组件
	casbin.Setup()

	r := gin.New()
	//加载路由
	server.Init(r)
	//注册中间件
	//	拦截器

	r.Use()
	//服务器启动
	if err := r.Run(":8000"); err != nil {
		fmt.Printf("Server start failed, err:%v\n", err)
	}
}
