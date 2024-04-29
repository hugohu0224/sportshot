package initinal

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "sportshot/proto"
	"sportshot/webservice/global"
)

func InitEventServerConn() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to connect: %v", err)
	}

	global.EnevtServerClient = pb.NewEventServiceClient(conn)
}
