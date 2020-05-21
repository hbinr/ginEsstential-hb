//File  : routes.go
//Author: duanhaobin
//Date  : 2020/5/20

package router

import (
	"ginEssential-hb/controller"
	"ginEssential-hb/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/auth/register", controller.Register)
		apiGroup.POST("/auth/login", controller.Login)
		// 用 认证中间件 保护用户信息接口
		apiGroup.POST("/auth/info", middleware.AuthMiddleWare(), controller.Info)
	}
	return r
}
