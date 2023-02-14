package access_control

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		ug := v1.Group("/ac")
		{
			ug.GET("/permission", getPermission)
			ug.PUT("/permission", putPermission)
			ug.POST("/permission", postPermission)
			ug.DELETE("/permission", deletePermission)
		}
	}
}
