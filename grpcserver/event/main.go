package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"sportshot/grpcserver/event/service"
	"sportshot/utils/config"
	"sportshot/utils/db"
	"sportshot/utils/global"
	pb "sportshot/utils/proto"
	"sportshot/utils/tools"
	"time"
)

func main() {
	// initial logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Panic(err.Error())
	}
	zap.ReplaceGlobals(logger)

	// initial config
	config.InitConfigByViper()
	uri := db.GetMongodbURI()
	zap.S().Infof("viper(config system) initialized")

	// initial MongodbClient
	global.MongodbClient = db.GetMongodbClient(uri)
	defer func(c *mongo.Client, ctx context.Context) {
		err := c.Disconnect(ctx)
		if err != nil {
			zap.S().Fatal("error disconnecting from mongodb:", err)
		}
	}(global.MongodbClient, context.TODO())
	zap.S().Infof("mongoClient initialized")

	// initial server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to listen: %v", err), zap.Error(err))
	}

	// getting IP that may change due to reboots
	localIP, err := tools.GetLocalIP()
	if err != nil {
		log.Fatalf("failed to get local IP: %v", err)
	}

	// creating an etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%s", viper.GetString("etcd.host"), viper.GetString("etcd.port"))},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer cli.Close()

	// register the service to etcd
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serviceKey := "myservice/" + localIP + ":8080" // 服务的键，这里包括IP和端口
	serviceValue := localIP
	_, err = cli.Put(ctx, serviceKey, serviceValue)
	if err != nil {
		log.Fatalf("Failed to set etcd key: %v", err)
	}
	log.Printf("Service registered with IP: %s", localIP)

	// start to serve
	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &service.EventServer{})
	zap.S().Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to serve: %v", err), zap.Error(err))
	}

}
