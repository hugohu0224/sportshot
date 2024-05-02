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

func (s *eventServer) SearchEvents(ctx context.Context, req *event.SearchEventsRequest) (*event.EventsReply, error) {
	searchFilter := bson.M{}

	// filter setting by request arguments
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

	// start to search
	zap.S().Debugf("search filter: %v", searchFilter)
	zap.S().Info("start to search")
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
	var data event.EventInfo
	reply := &event.EventsReply{
		EventInfo: nil,
		Message:   "query success",
		Status:    200,
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	// fetch single events document
	for _, result := range results {
		// get events
		if pvs, ok := result["events"].(primitive.A); ok {
			// get single event
			for _, pv := range pvs {
				mpv, err := bson.Marshal(pv)
				if err != nil {
					return nil, err
				}
				if err := bson.Unmarshal(mpv, &data); err != nil {
					return nil, err
				}
				// append single event to EventInfo
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
