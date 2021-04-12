package main

import (
	"cnhtc/gin-vue-admin/AppServer/core"
	"cnhtc/gin-vue-admin/AppServer/global"
	"cnhtc/gin-vue-admin/AppServer/initialize"
	"fmt"
)

/**
http://www.topgoer.com 学习网站
*/
func main() {
	//初始化全局配置
	global.GVA_VP = core.Viper()
	global.GVA_DB = initialize.Sqlx()
	fmt.Println(global.GVA_CONFIG.System.Addr)
	core.RunWindowsServer()
}
