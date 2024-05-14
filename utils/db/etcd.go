package db

import (
	"fmt"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"time"
)

func discoverService(serviceKey string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%s", viper.GetString("etcd.host"), viper.GetString("etcd.port"))},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		zap.S().Errorf("error connecting to etcd: %v", err)
	}

}
