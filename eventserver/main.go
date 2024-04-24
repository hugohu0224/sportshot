package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	pb "sportshot/proto"
)

type eventServer struct {
	pb.UnimplementedEventServiceServer
}

func (s *eventServer) SearchEvents(ctx context.Context, req *pb.SearchEventsRequest) (*pb.EventsReply, error) {
	file, err := os.Open("events.json")
	if err != nil {
		fmt.Println("無法打開文件:", err)
		return &pb.EventsReply{}, err
	}
	defer file.Close()

	// 創建一個新的Decoder來解析JSON
	decoder := json.NewDecoder(file)

	// 解析JSON並將數據存儲到result中
	result := &pb.EventsReply{}
	err = decoder.Decode(&result.EventInfo)
	if err != nil {
		fmt.Println("解析JSON時發生錯誤:", err)
		return &pb.EventsReply{}, err
	}

	return result, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &eventServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
