package auth

import (
	"casbin-gin/common/component/casbin"
	"casbin-gin/common/component/jwtauth"
	"casbin-gin/common/component/response"
	"casbin-gin/common/middleware"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Authorize 用户授权：确定用户身份后，通过权限策略对用户权限校验及授权（检查通过就算是授权了）
// 通过JWT payload 带用户身份信息过来
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		//log := api.GetRequestLogger(c)
		data, _ := c.Get(jwtauth.JwtPayloadKey)
		v := data.(jwtauth.MapClaims)
		//e := sdk.Runtime.GetCasbinKey(c.Request.Host)
		e := casbin.Enforcer
		var res, casbinExclude bool
		var err error

		//检查权限
		//1 如果是管理员直接通过
		if v["rolekey"] == "admin" {
			res = true
			c.Next()
			return
		}
		//2 接口是否需要权限检查，有些不需要权限检查直接通过
		for _, i := range middleware.CasbinExclude {
			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
				casbinExclude = true
				break
			}
		}
		if casbinExclude {
			log.Printf("Casbin exclusion, no validation method:%s path:%s\n", c.Request.Method, c.Request.URL.Path)
			c.Next()
			return
		}
		//3 调用Casbin的权限校验方法
		res, err = e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		if err != nil {
			log.Printf("AuthCheckRole error:%s method:%s path:%s\n", err, c.Request.Method, c.Request.URL.Path)
			response.Error(c, 500, err, "")
			return
		}

		if res { //校验通过交给下一个责任链节点处理
			log.Printf("isTrue: %v role: %s method: %s path: %s\n", res, v["rolekey"], c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else { //校验失败，直接响应结果
			log.Printf("isTrue: %v role: %s method: %s path: %s message: %s\n",
				res, v["rolekey"], c.Request.Method, c.Request.URL.Path, "当前request无权限，请管理员确认！")
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}
	}
}
