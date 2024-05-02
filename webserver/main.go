package main

import (
	"go.uber.org/zap"
	"sportshot/webserver/initinal"
)

func main() {
	// initialize
	initinal.InitLogger()
	zap.S().Info("logger initialized")
	initinal.InitEventServerConn()
	zap.S().Info("event server initialized")
	Router := initinal.InitRouters()
	zap.S().Info("router initialized")

	err := Router.Run("localhost:8081")
	if err != nil {
		zap.S().Panicf("fail to start web server")
	}
}
