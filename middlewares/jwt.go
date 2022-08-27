package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lily-lee/blog-server/services/jwttoken"
)

func JwtAuth(skip ...bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var isSkip bool
		if len(skip) > 0 {
			isSkip = skip[0]
		}

		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			if isSkip {
				c.Next()
				return
			}

			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		claims, err := jwttoken.ParseToken(tokenStr)
		if err != nil {
			if isSkip {
				c.Next()
				return
			}

			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("user_claims", claims)
		c.Next()
	}
}
