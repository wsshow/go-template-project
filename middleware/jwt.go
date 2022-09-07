package middleware

import (
	"github.com/gin-gonic/gin"
	"gtp/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(200, gin.H{"code": 101, "msg": "token is null"})
			c.Abort()
			return
		}
		j := utils.NewJWT()
		_, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				c.JSON(200, gin.H{"code": 102, "msg": "token is expired"})
				c.Abort()
				return
			}
			c.JSON(200, gin.H{"code": 103, "msg": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
