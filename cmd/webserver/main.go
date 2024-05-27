package main

import (
	"go.uber.org/zap"
	"sportshot/internal/webserver/initinal"
	"sportshot/pkg/utils/db"
	"sportshot/pkg/utils/global"
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
	err := Router.Run("0.0.0.0:8080")
	if err != nil {
		zap.S().Panicf("fail to start web server")
	}
}
