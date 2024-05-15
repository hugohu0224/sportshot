package db

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.uber.org/zap"
	"sportshot/utils/global"
)

func GetEtcdClient() *clientv3.Client {
	etcdURL := fmt.Sprintf("%s:%d", viper.GetString("etcd.host"), viper.GetInt("etcd.port"))
	cli, err := clientv3.NewFromURL(etcdURL)
	if err != nil {
		zap.S().Fatalf("failed to connect to etcd: %v", err)
		return nil
	}
	return cli
}

func RegisterServiceToEtcd(cli *clientv3.Client, serverName string, serverHost string, serverPort string) error {
	serverValue := fmt.Sprintf("%s:%s", serverHost, serverPort)
	serverKey := fmt.Sprintf("%s/%s", serverName, serverValue)

	em, err := endpoints.NewManager(global.EtcdClient, fmt.Sprintf("%s", serverName))
	if err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to create endpoint manager: %v", err), zap.Error(err))
	}
	err = em.AddEndpoint(context.TODO(), serverKey, endpoints.Endpoint{Addr: serverValue})
	if err != nil {
		zap.S().Fatalf("failed to register to etcd: %v", err)
		return err
	}
	zap.S().Infof("server registered to etcd by KEY: %s", serverKey)

	return nil
}
