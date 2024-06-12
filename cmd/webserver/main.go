package main

import (
	"fmt"
	"go.uber.org/zap"
	"sportshot/internal/webserver/initinal"
	"sportshot/pkg/utils/db"
	"sportshot/pkg/utils/global"
	"sportshot/pkg/utils/tools"
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

	// mysql
	global.MySQLClient = db.GetMySQLClient()
	zap.S().Info("mysql client initialized")

	// shortcut to visit webserver
	host, err := tools.GetLocalHost()
	if err != nil {
		zap.S().Errorf("get local host err: %v", err)
	}
	fmt.Printf("Click here to visit SportShot --->  http://%v:8080/v1/auth/login \n", host)

	// start
	err = Router.Run(":8080")
	if err != nil {
		zap.S().Panicf("fail to start web server")
	}
}
