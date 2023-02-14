package auth

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Authorize intercepted ...")

	}
}
