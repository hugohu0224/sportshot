package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sportshot/internal/grpcserver/event/service"
	db2 "sportshot/pkg/utils/db"
	"sportshot/pkg/utils/global"
	pb "sportshot/pkg/utils/proto"
	"sportshot/pkg/utils/tools"
)

// local config consider to move to viper
var (
	serverName    = "event"
	serverPort    = "50051"
	serverHost, _ = tools.GetLocalHost()
)

func main() {
	// logger
	logger, err := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	// config
	db2.InitConfigByViper()
	zap.S().Infof("viper initialized")

	// mongodb
	global.MongodbClient = db2.GetMongodbClient()
	defer global.MongodbClient.Disconnect(context.TODO())
	zap.S().Infof("mongoClient initialized")

	// etcd
	global.EtcdClient = db2.GetEtcdClient()
	defer global.EtcdClient.Close()
	zap.S().Info("etcd client initialized")

	// register
	err = db2.RegisterServiceToEtcd(global.EtcdClient, serverName, serverHost, serverPort)
	if err != nil {
		zap.S().Fatalf("register service to etcd failed: %v", err)
	}

	// start to serve
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to listen: %v", err), zap.Error(err))
	}
	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &service.EventServer{})
	zap.S().Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to serve: %v", err), zap.Error(err))
	}
}
