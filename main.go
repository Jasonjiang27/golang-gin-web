package main

import (
	"fmt"
	"log"
	"golang-gin-web/models"
	"golang-gin-web/pkg/setting"
	"golang-gin-web/routers"
	"net/http"
	//"github.com/fvbock/endless"  热更新应用在windows使用失败，安装报错
)

func main() {
	setting.Setup()
	models.Setup()

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
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("server err :%v", err)
	}

}
