package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"sportshot/proto"
)

type eventServer struct {
	event.UnimplementedEventServiceServer
}

func (s *eventServer) SearchEvents(ctx context.Context, req *event.SearchEventsRequest) (*event.EventsReply, error) {
	fileName := "events.json"
	file, err := os.Open(fileName)
	if err != nil {
		zap.S().Info(fmt.Sprintf("Error opening file %s", fileName), zap.Error(err))
		return &event.EventsReply{}, err
	}
	defer file.Close()

	// 創建一個新的Decoder來解析JSON
	decoder := json.NewDecoder(file)

	// 解析JSON並將數據存儲到result中
	result := &event.EventsReply{}
	err = decoder.Decode(&result.EventInfo)
	if err != nil {
		zap.S().Info(fmt.Sprintf("Error parsing json file %s", fileName), zap.Error(err))
		return &event.EventsReply{}, err
	}

	return result, nil
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
