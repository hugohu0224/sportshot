package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sportshot/grpcserver/event/service"
	"sportshot/utils/db"
	"sportshot/utils/global"
	pb "sportshot/utils/proto"
	"sportshot/utils/tools"
)

// local config consider to move to viper
const (
	serviceName = "/event"
	servicePort = "50051"
)

func main() {
	// initial logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Panic(err.Error())
	}
	zap.ReplaceGlobals(logger)

	// initial config
	db.InitConfigByViper()
	uri := db.GetMongodbURI()
	zap.S().Infof("viper(config system) initialized")

	// initial MongodbClient
	global.MongodbClient = db.GetMongodbClient(uri)
	defer func(ctx context.Context, c *mongo.Client) {
		err := c.Disconnect(ctx)
		if err != nil {
			zap.S().Fatal("error disconnecting from mongodb:", err)
		}
	}(context.TODO(), global.MongodbClient)
	zap.S().Infof("mongoClient initialized")

	// initial server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", servicePort))
	if err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to listen: %v", err), zap.Error(err))
	}

	// getting IP that may change due to reboots
	localHost, err := tools.GetLocalHost()
	if err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to get local ip: %v", err), zap.Error(err))
	}

	// getting etcd client
	cli, err := db.GetEtcdClient()
	if err != nil {
		zap.S().Errorf("fail to GetEtcdClient : %v", err)
	}
	defer func(cli *clientv3.Client) {
		err := cli.Close()
		if err != nil {
			zap.S().Errorf("error closing etcd: %v", err)
		}
	}(cli)

	// register the service to etcd
	serviceAddr := fmt.Sprintf("%s:%s", localHost, servicePort)
	serviceKey := fmt.Sprintf("%s/%s", serviceName, "server1")
	err = db.RegisterToEtcd(cli, serviceKey, serviceAddr)
	if err != nil {
		zap.S().Fatalf("failed to register to etcd: %v", err)
	}
	zap.S().Infof("server registered to etcd by KEY: %s VALUE: %s", serviceKey, serviceAddr)

	// start to serve
	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &service.EventServer{})
	zap.S().Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to serve: %v", err), zap.Error(err))
	}

}
