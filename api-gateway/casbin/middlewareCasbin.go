package rule

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// CasbinMiddleware - Casbin ruxsatlarini tekshirish uchun middleware
func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-User-Role")
		if role == "" {
			role = "user" // default role
		}

		// Casbin ruxsatlarni tekshirish
		ok, _ := enforcer.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access Denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
