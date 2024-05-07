package main

import (
	"go.uber.org/zap"
	"sportshot/webserver/initinal"
)

func main() {
	// Logger
	initinal.InitLogger()
	zap.S().Info("logger initialized")

	// EventServer
	initinal.InitEventServerConn()
	zap.S().Info("eventserver initialized")

	// Routers
	Router := initinal.InitRouters()
	zap.S().Info("router initialized")

	// Run
	err := Router.Run("localhost:8081")
	if err != nil {
		zap.S().Panicf("fail to start web server")
	}
}
