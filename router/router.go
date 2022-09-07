package router

import (
	"gtp/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	rpcap := r.Group("/gtp")
	{
		// 获取当前版本
		rpcap.GET("/devicesName", controller.GetVersion())
	}
}
