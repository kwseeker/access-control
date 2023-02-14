package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

/*
用户相关请求处理
1）注册
2）登录
3）查询用户列表
4）用户权限修改
*/

func registerHandler(ctx *gin.Context) {
	log.Println("call registerHandler() ...")
	//TODO
	ctx.String(200, fmt.Sprintf("register ..."))
}

func loginHandler(ctx *gin.Context) {
	log.Println("call loginHandler() ...")
	//TODO
	ctx.String(200, fmt.Sprintf("login ..."))
}
