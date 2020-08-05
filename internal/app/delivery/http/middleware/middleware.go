package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skazak/example/internal/app/delivery/http/auth"
	"github.com/skazak/example/internal/app/delivery/http/response"
)

func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Auth-Token")
		if token == "" {
			c.JSON(http.StatusForbidden, response.ErrorResponse{
				Message: "token is not provided",
			})
			c.Abort()
			return
		}

		claims, err := auth.DecodeClaims(token)
		if err != nil {
			c.JSON(http.StatusForbidden, response.ErrorResponse{
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		if claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusForbidden, response.ErrorResponse{
				Message: "token expired",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
