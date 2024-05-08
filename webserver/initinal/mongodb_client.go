package initinal

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"sportshot/utils/global"
)

func InitMongoClient(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		zap.S().Fatalf("Failed to connect to MongoDB : %v", err)
	}
	// 測試連線
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		zap.S().Fatalf("Failed to ping MongoDB : %v", err)
	}
	global.MongodbClient = client
}
