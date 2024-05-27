package db

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"os"
)

func GetEtcdClient() *clientv3.Client {
	etcdURL := os.Getenv("ETCD_CONN")
	zap.S().Infof("connecting to etcd at %s", etcdURL)

	cli, err := clientv3.NewFromURL(etcdURL)
	if err != nil {
		zap.S().Fatalf("failed to connect to etcd: %v", err)
		return nil
	}
	return cli
}

func RegisterServiceToEtcd(cli *clientv3.Client, serverName string, serverHost string, serverPort string) error {
	// lease
	lease, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		zap.S().Fatalf("failed to create lease: %v", err)
		return err
	}

	// settings
	serverInfo := map[string]string{
		"key":   fmt.Sprintf("%s/%s:%s", serverName, serverHost, serverPort),
		"value": fmt.Sprintf("{\"Addr\":\"%s:%s\", \"LeaseId\":\"%d\"}", serverHost, serverPort, lease.ID),
	}

	// register
	_, err = cli.Put(context.TODO(), serverInfo["key"], serverInfo["value"], clientv3.WithLease(lease.ID))
	if err != nil {
		zap.S().Errorf("Failed to register to etcd with lease: %v", err)
		return err
	}

	// keep
	ch, err := cli.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		zap.S().Errorf("failed to keep lease alive: %v", err)
		return err
	}

	// consume ch to avoid full message
	go func() {
		for ka := range ch {
			if ka == nil {
				zap.S().Info("keepalive channel closed")
				return
			}
			// keepalive nothing to do
		}
	}()

	zap.S().Infof("server registered to etcd by KEY: %s", serverInfo["value"])

	return nil
}
