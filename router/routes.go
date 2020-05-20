//File  : routes.go
//Author: duanhaobin
//Date  : 2020/5/20

package router

import (
	"ginEssential-hb/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/auth/register", controller.Register)
		apiGroup.POST("/auth/login", controller.Login)
	}
	return r
}
