package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"sportshot/proto"
)

type eventServer struct {
	event.UnimplementedEventServiceServer
}

// SearchEvents localize search for testing
//func (s *eventServer) SearchEvents(ctx context.Context, req *event.SearchEventsRequest) (*event.EventsReply, error) {
//	fileName := "events.json"
//	file, err := os.Open(fileName)
//	if err != nil {
//		zap.S().Info(fmt.Sprintf("Error opening file %s", fileName), zap.Error(err))
//		return &event.EventsReply{}, err
//	}
//
//	defer func(file *os.File) {
//		err := file.Close()
//		if err != nil {
//			zap.S().Info(fmt.Sprintf("Error closing file %s", file.Name()), zap.Error(err))
//		}
//	}(file)
//
//	decoder := json.NewDecoder(file)
//	result := &event.EventsReply{}
//	err = decoder.Decode(&result.EventInfo)
//	if err != nil {
//		zap.S().Info(fmt.Sprintf("Error parsing json file %s", fileName), zap.Error(err))
//		return &event.EventsReply{}, err
//	}
//
//	return result, nil
//}

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
	// implement mongodb query here
	//collection := db.Collection("event")
	//cursor, err := collection.Find(context.TODO(), filter)
	//if err != nil {
	//	return nil, err
	//}
	//defer cursor.Close(context.TODO())
	//
	//var results []bson.M
	//if err = cursor.All(context.TODO(), &results); err != nil {
	//	return nil, err
	//}

	return &event.EventsReply{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to listen: %v", err), zap.Error(err))
	}
	s := grpc.NewServer()
	event.RegisterEventServiceServer(s, &eventServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		zap.S().Panicf(fmt.Sprintf("failed to serve: %v", err), zap.Error(err))
	}
}
