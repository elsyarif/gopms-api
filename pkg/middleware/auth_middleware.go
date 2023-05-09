package middleware

import (
	"github.com/elsyarif/pms-api/pkg/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ResponseJSON.Error("fail", "Unauthorized", nil))
			return
		}

		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ResponseJSON.Error("fail", "invalid token", nil))
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		claims, err := helper.VerifyToken(tokenString, helper.AccessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ResponseJSON.Error("fail", "Unauthorized", nil))
			return
		}

		c.Set("users", claims)
		c.Next()
	}
}
