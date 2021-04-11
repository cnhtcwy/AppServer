package router

import (
	v1 "cnhtc/gin-vue-admin/AppServer/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitBaseRouter(Router *gin.RouterGroup)(R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login",v1.Login)
		BaseRouter.POST("captcha",v1.Captcha)
		BaseRouter.GET("public", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"public":"OK",
			})
		})
	}
	return BaseRouter
}
