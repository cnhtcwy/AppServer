package core

import (
	"cnhtc/gin-vue-admin/AppServer/initialize"
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunWindowsServer()  {
	fmt.Println("http服务启动")
	router := initialize.Routers()
	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}


//func RunOtherServer(address string, router *gin.Engine) server {
//	s := endless.NewServer(address, router)
//	s.ReadHeaderTimeout = 10 * time.Millisecond
//	s.WriteTimeout = 10 * time.Second
//	s.MaxHeaderBytes = 1 << 20
//	return s
//}