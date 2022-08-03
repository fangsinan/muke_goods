package middlewares

import (
	"net/http"
	"webApi/user_web/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		currentUser := claims.(*models.CustomClaims)
		zap.S().Infof("[isAdmin] currentUser", currentUser.AuthorityId)
		if currentUser.AuthorityId == 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
