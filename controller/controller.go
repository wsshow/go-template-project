package controller

import (
	"github.com/gin-gonic/gin"
	"gtp/utils"
	"gtp/version"
	"net/http"
)

var resp utils.Response

func GetVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, resp.Success(version.Get))
	}
}
