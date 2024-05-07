package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sportshot/utils/config"
	"sportshot/utils/db"
	pb "sportshot/utils/proto"
	"sportshot/webserver/global"
	"time"
)

type eventServer struct {
	pb.UnimplementedEventServiceServer
}

func (s *eventServer) SearchEvents(ctx context.Context, req *pb.SearchEventsRequest) (*pb.EventsReply, error) {
	// initial filter
	filter := bson.D{}

	// default filter
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	filter = append(filter, bson.E{Key: "date", Value: bson.M{"$gte": yesterday}})

	// condition filter
	if req.LeagueName != "" {
		filter = append(filter, bson.E{Key: "leagueName", Value: req.LeagueName})
	}
	if req.SportType != "" {
		filter = append(filter, bson.E{Key: "sportType", Value: req.SportType})
	}
	if req.StartDate != "" {
		filter = append(filter, bson.E{Key: "date", Value: bson.M{"$gte": req.StartDate}})
	}
	if req.EndDate != "" {
		filter = append(filter, bson.E{Key: "date", Value: bson.M{"$lte": req.EndDate}})
	}

	// connect to mongodb
	databaseName := "sportevents"
	collectionName := "basketball"
	collection := global.MongodbClient.Database(databaseName).Collection(collectionName)
	zap.S().Infof("connected to mongodb, start to search events")

	// start to search
	zap.S().Infof("start to search by filter : %v", filter)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	zap.S().Infof("search successfully")
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			zap.S().Warnf("failed to close cursor, %v", err)
		}
	}(cursor, context.TODO())

	// processing data
	var results []*pb.EventInfo
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	// reorg reply
	reply := &pb.EventsReply{
		Events:  results,
		Message: "query success",
		Status:  200,
	}

	return reply, nil
}
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

	// start to serve
	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &eventServer{})
	zap.S().Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to serve: %v", err), zap.Error(err))
	}

}
