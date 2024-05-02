package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"sportshot/utils/db"
	"sportshot/utils/proto"
	"sportshot/webserver/global"
)

type eventServer struct {
	event.UnimplementedEventServiceServer
}

func getString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

func getInt64(v interface{}) int64 {
	if i, ok := v.(int64); ok {
		return i
	}
	return 0
}

func (s *eventServer) SearchEvents(ctx context.Context, req *event.SearchEventsRequest) (*event.EventsReply, error) {
	searchFilter := bson.M{}

	if req.Name != "" {
		searchFilter["name"] = req.Name
	}
	if req.Type != "" {
		searchFilter["type"] = req.Type
	}
	if req.Date != "" {
		searchFilter["date"] = req.Date
	}

	// connect to mongodb
	databaseName := "sportevents"
	collectionName := "basketball"
	collection := global.MongodbClient.Database(databaseName).Collection(collectionName)
	zap.S().Infof("connected to mongodb, start to search events")
	cursor, err := collection.Find(context.TODO(), searchFilter)
	if err != nil {
		return nil, err
	}
	zap.S().Infof("search events successfully")
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			zap.S().Warnf("failed to close cursor, %v", err)
		}
	}(cursor, context.TODO())

	// processing data
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	// data from bson.M to EventsReply
	reply := &event.EventsReply{
		EventInfo: nil,
		Message:   "query success",
		Status:    200,
	}

	for _, result := range results {
		if pvs, ok := result["events"].(primitive.A); ok {
			var data event.EventInfo
			for _, pv := range pvs {
				mpv, err := bson.Marshal(pv)
				if err != nil {
					return nil, err
				}
				if err := bson.Unmarshal(mpv, &data); err != nil {
					return nil, err
				}
				reply.EventInfo = append(reply.EventInfo, &data)
			}
		}

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

	// read config
	zap.S().Infof("reading config file...")
	data, err := os.ReadFile("config.json")
	if err != nil {
		zap.S().Fatalf("error reading config.json: %v", err)
	}
	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		zap.S().Fatalf("error parsing config.json: %v", err)
	}
	zap.S().Infof("loaded config %v", config)

	// initial mongodb
	uri := config["mongodbURI"].(string)
	global.MongodbClient = db.GetMongoClient(uri)

	// initial server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to listen: %v", err), zap.Error(err))
	}
	s := grpc.NewServer()
	event.RegisterEventServiceServer(s, &eventServer{})
	zap.S().Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to serve: %v", err), zap.Error(err))
	}

}
