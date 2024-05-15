package main

import (
	"go.uber.org/zap"
	"sportshot/utils/db"
	"sportshot/utils/global"
	"sportshot/webserver/initinal"
)

func main() {
	// logger
	initinal.InitLogger()
	zap.S().Info("logger initialized")

	// viper
	db.InitConfigByViper()
	zap.S().Infof("viper initialized")

	// routers
	Router := initinal.InitRouters()
	zap.S().Info("router initialized")

	// etcd
	global.EtcdClient = db.GetEtcdClient()
	defer global.EtcdClient.Close()
	zap.S().Info("etcd client initialized")

	// start
	err := Router.Run("localhost:8081")
	if err != nil {
		zap.S().Panicf("fail to start web server")
	}
}
