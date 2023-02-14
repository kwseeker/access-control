package server

import (
	"casbin-gin/app/admin/routers/access_control"
	"casbin-gin/app/admin/routers/user"
	"github.com/gin-gonic/gin"
)

type Option func(engine *gin.Engine)

var options []Option

// Register 注册路由
func Register(opts ...Option) {
	options = append(options, opts...)
}

func RegisterDefault() {
	Register(user.Routers, access_control.Routers)
}

// Init 初始化gin引擎并加载路由
func Init(r *gin.Engine) {
	RegisterDefault()
	for _, opt := range options {
		opt(r)
	}
}
