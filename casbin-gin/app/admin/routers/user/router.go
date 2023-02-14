package user

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		ug := v1.Group("/user")
		{
			ug.PUT("/register", registerHandler)
			ug.GET("/login", loginHandler)
		}
	}
}
