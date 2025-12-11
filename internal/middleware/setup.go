package middleware

import "github.com/gin-gonic/gin"

func SetupMiddlewares(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(CORS())
	router.Use(HandlerError())
}
