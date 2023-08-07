package middlewares

import (
	"go-bankmate/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := util.TokenValid(ctx)
		if err != nil {
			log.Println(err)
			ctx.String(http.StatusUnauthorized, "unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
