package middleware

import (
	"errors"
	"go_template/internal/utils/apperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		last := c.Errors.Last()

		var appErr apperror.AppError
		if errors.As(last.Err, &appErr) {
			c.AbortWithStatusJSON(appErr.Code, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"status":  appErr.Status,
			})
			return
		}

		// fallback genérico a 500 si no es un AppError tipado
		fallback := apperror.FromCode(http.StatusInternalServerError, last.Err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    fallback.Code,
			"message": fallback.Message,
			"status":  fallback.Status,
		})
	}
}
