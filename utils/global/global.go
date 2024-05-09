package global

import (
	"go.mongodb.org/mongo-driver/mongo"
	pb "sportshot/utils/proto"
)

var (
	EventServerClient pb.EventServiceClient
	MongodbClient     *mongo.Client
)