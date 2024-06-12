package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sportshot/internal/grpcserver/events"
	"sportshot/pkg/utils/db"
	"sportshot/pkg/utils/global"
	pb "sportshot/pkg/utils/proto"
	"sportshot/pkg/utils/tools"
)

// local config consider to move to viper
var (
	serverName    = "events"
	serverPort    = "50051"
	serverHost, _ = tools.GetLocalHost()
)

func main() {
	// logger
	logger, err := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	// config
	db.InitConfigByViper()
	zap.S().Infof("viper initialized")

	// mongodb
	global.MongodbClient = db.GetMongodbClient()
	defer global.MongodbClient.Disconnect(context.TODO())
	zap.S().Infof("mongoClient initialized")

	// etcd
	global.EtcdClient = db.GetEtcdClient()
	defer global.EtcdClient.Close()
	zap.S().Info("etcd client initialized")

	// register
	err = db.RegisterServiceToEtcd(global.EtcdClient, serverName, serverHost, serverPort)
	if err != nil {
		zap.S().Fatalf("register service to etcd failed: %v", err)
	}

	// sometimes BetsAPI will ask users to log in during the rush time,
	// so the crawler can't crawl the data directly, in order to make the demo more convenient,
	// so we initial the old data in the beginning
	db.InitOldDataForDemo("sportevents", "basketball", "./pkg/files/sportevents.basketball.json")

	// start to serve
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to listen: %v", err), zap.Error(err))
	}
	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &events.EventServer{})
	zap.S().Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to serve: %v", err), zap.Error(err))
	}
}
