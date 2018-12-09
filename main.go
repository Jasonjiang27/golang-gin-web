package main

import (
	"fmt"
	"golang-gin-web/models"
	"golang-gin-web/pkg/setting"
	"golang-gin-web/routers"
	"net/http"
	//"github.com/fvbock/endless"  热更新应用在windows使用失败，安装报错
)

func main() {
	setting.Setup()
	models.Setup()
	router := routers.InitRouter()
	/*
			endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
			endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
			endless.DefaultMaxHeaderBytes = 1 << 20


		DefaultReadTimeOut := setting.ServerSetting.ReadTimeout
		DefaultWriteTimeOut := setting.ServerSetting.WriteTimeout
		DefaultMaxHeaderBytes := 1 << 20
		endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

		server := endless.NewServer(endPoint, routers.InitRouter())
		server := NewServer(endPoint, routers.InitRouter())
		server.BeforeBegin = func(add string) {
			log.Printf("Actual pid is %d", syscall.Getpid())
		}

		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server err: %v", err)
		}
	*/
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
