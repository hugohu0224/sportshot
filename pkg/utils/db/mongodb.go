package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"os"
)

func GetMongodbClient() *mongo.Client {
	uri := getMongodbURI()
	zap.S().Infof("connecting to MongoDB at %s", uri)
	clientOptions := options.Client().ApplyURI(uri).
		SetMaxPoolSize(20).
		SetMinPoolSize(5).
		SetMaxConnIdleTime(30)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		zap.S().Fatalf("failed to connect to MongoDB : %v", err)
	}

	// test connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		zap.S().Fatalf("failed to ping MongoDB : %v", err)
	}
	return client
}

func getMongodbURI() string {
	//user := viper.GetString("mongodb.user")
	//password := viper.GetString("mongodb.password")
	//host := viper.GetString("mongodb.host")
	//port := viper.GetString("mongodb.port")
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)

	uri := os.Getenv("MONGODB_CONN")
	zap.S().Infof("connecting to MongoDB at %s", uri)
	return uri
}