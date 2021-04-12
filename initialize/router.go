package initialize

import (
	"cnhtc/gin-vue-admin/AppServer/global"
	"cnhtc/gin-vue-admin/AppServer/middleware"
	"cnhtc/gin-vue-admin/AppServer/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	fmt.Println("路由初始化开始")
	Router := gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	//}
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//路由分组
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//router.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitApiRouter(PrivateGroup)                   // 注册功能api路由
		router.InitJwtRouter(PrivateGroup)                   // jwt相关路由
		router.InitUserRouter(PrivateGroup)                  // 注册用户路由
		router.InitMenuRouter(PrivateGroup)                  // 注册menu路由
		router.InitEmailRouter(PrivateGroup)                 // 邮件相关路由
		router.InitSystemRouter(PrivateGroup)                // system相关路由
		router.InitCasbinRouter(PrivateGroup)                // 权限相关路由
		router.InitCustomerRouter(PrivateGroup)              // 客户路由
		router.InitAutoCodeRouter(PrivateGroup)              // 创建自动化代码
		router.InitAuthorityRouter(PrivateGroup)             // 注册角色路由
		router.InitSimpleUploaderRouter(PrivateGroup)        // 断点续传（插件版）
		router.InitSysDictionaryRouter(PrivateGroup)         // 字典管理
		router.InitSysOperationRecordRouter(PrivateGroup)    // 操作记录
		router.InitSysDictionaryDetailRouter(PrivateGroup)   // 字典详情管理
		router.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
		router.InitExcelRouter(PrivateGroup)                 // 表格导入导出
		PrivateGroup.GET("private", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"private": "OK",
			})
		})
	}
	fmt.Println("路由初始化结束")
	return Router
}
