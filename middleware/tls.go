package middleware

import (
	"gtp/log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func LoadTls(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			log.Error("loadTls error:", err)
			return
		}
		c.Next()
	}
}
