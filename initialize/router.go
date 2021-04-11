package initialize

import (
	"cnhtc/gin-vue-admin/AppServer/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Routers() *gin.Engine {
	fmt.Println("路由初始化开始")
	Router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	//路由分组
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		PrivateGroup.GET("private", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"private":"OK",
			})
		})
	}
	fmt.Println("路由初始化结束")
	return Router
}
