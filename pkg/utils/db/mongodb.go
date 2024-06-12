package db

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
	"os"
	"sportshot/pkg/utils/global"
	"sportshot/pkg/utils/models"
	"sportshot/pkg/utils/tools"
)

func GetMongodbClient() *mongo.Client {
	uri := os.Getenv("MONGODB_CONN")
	zap.S().Infof("connecting to MongoDB at %s", uri)

	clientOptions := options.Client().ApplyURI(uri).
		SetMaxPoolSize(20).
		SetMinPoolSize(5).
		SetMaxConnIdleTime(30)

	// connect to mongodb
	ctx, cancel := tools.TimeOutCtx(10)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		zap.S().Fatalf("failed to connect to MongoDB : %v", err)
	}

	// test connection
	ctx, cancel = tools.TimeOutCtx(3)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		zap.S().Fatalf("failed to ping MongoDB : %v", err)
	}
	return client
}

func InitOldDataForDemo(dbName string, collectionName string, filePath string) {

	database := global.MongodbClient.Database(dbName)
	collection := database.Collection(collectionName)

	// check if any
	ctxAny, cancelAny := tools.TimeOutCtx(5)
	defer cancelAny()
	err := collection.FindOne(ctxAny, bson.M{}).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			zap.S().Info("no any data in mongodb, start to initial")
		} else {
			zap.S().Fatalf("failed initial data to mongodb : %v", err)
		}
	} else {
		zap.S().Info("already had data in mongodb, no need to initial")
		return
	}

	// read
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal
	var events []models.SportEvent
	err = json.Unmarshal(jsonData, &events)
	if err != nil {
		log.Fatal(err)
	}

	// cast to interface{}
	var docs []interface{}
	for _, e := range events {
		docs = append(docs, e)
	}

	// insert
	ctxInsert, cancelInsert := tools.TimeOutCtx(5)
	defer cancelInsert()
	_, err = collection.InsertMany(ctxInsert, docs)
	if err != nil {
		zap.S().Errorf("failed to insert data into mongodb : %v", err)
	}
	zap.S().Info("successfully inserted data into mongodb")
}
