package main

import (
	"go.uber.org/zap"
	"sportshot/webserver/initinal"
)

func main() {
	// initialize
	initinal.InitLogger()
	zap.S().Info("Logger initialized")
	initinal.InitEventServerConn()
	zap.S().Info("Event server initialized")
	Router := initinal.InitRouters()
	zap.S().Info("Router initialized")

	err := Router.Run("localhost:8081")
	if err != nil {
		zap.S().Panicf("fail to start web server")
	}
}
