package access_control

import (
	"casbin-gin/common/component/casbin"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

/*
Casbin 如果使用关系型数据库存储
数据表结构
 	id	ptype	v0	v1	v2	v3	v4	v5
*/

// 查询策略
func getPermission(ctx *gin.Context) {
	//TODO
	ctx.String(200, fmt.Sprintf("getPermission ..."))
}

// 添加策略
func postPermission(ctx *gin.Context) {
	//data, _ := ctx.GetRawData()
	//log.Println("Query raw data: ", data)
	ok, _ := casbin.Enforcer.AddPolicy("admin", "/api/v1/hello", "GET")
	if !ok {
		log.Println("policy already exist")
	} else {
		log.Println("add policy succeed")
	}
	ctx.String(200, fmt.Sprintf(strconv.FormatBool(ok)))
}

// 修改策略
func putPermission(ctx *gin.Context) {
	casbin.Enforcer.RemovePolicy()
	casbin.Enforcer.AddPolicy()
	ctx.String(200, fmt.Sprintf("putPermission ..."))
}

// 删除策略
func deletePermission(ctx *gin.Context) {
	if ok, _ := casbin.Enforcer.RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {

	}
	ctx.String(200, fmt.Sprintf("deletePermission ..."))
}
