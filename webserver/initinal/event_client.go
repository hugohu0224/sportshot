package initinal

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sportshot/utils/global"
	pb "sportshot/utils/proto"
)

func InitEventServerConn() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to connect to EventServer: %v", err)
	}

	global.EventServerClient = pb.NewEventServiceClient(conn)
}
