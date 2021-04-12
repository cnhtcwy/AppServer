package router

import (
	v1 "cnhtc/gin-vue-admin/AppServer/api/v1"
	"cnhtc/gin-vue-admin/AppServer/middleware"
	"github.com/gin-gonic/gin"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
