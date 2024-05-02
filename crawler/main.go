package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
	"sportshot/crawler/global"
	"sportshot/crawler/operator"
	"sportshot/utils/db"
	"time"
)

func main() {
	// initial logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infof("logger initialized")

	// initial basketball crawler
	bc := operator.BasketballCrawler{}
	zap.S().Infof("BasketballCrawler initialized")

	// read config
	data, err := os.ReadFile("config.json")
	if err != nil {
		zap.S().Fatalf("error reading config.json: %v", err)
	}
	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		zap.S().Fatalf("error parsing config.json: %v", err)
	}
	zap.S().Infof("loaded config %v", config)

	// initial MongoClient
	uri := config["mongodbURI"].(string)
	global.MongodbClient = db.GetMongoClient(uri)
	defer func(c *mongo.Client, ctx context.Context) {
		err := c.Disconnect(ctx)
		if err != nil {
			zap.S().Fatal("error disconnecting from mongodb:", err)
		}
	}(global.MongodbClient, context.TODO())
	zap.S().Infof("mongoClient initialized")

	// connect to mongo
	databaseName := "sportevents"
	collectionName := "basketball"
	collection := global.MongodbClient.Database(databaseName).Collection(collectionName)
	zap.S().Infof("Mongodb.%s.%s connected", databaseName, collectionName)

	// start to crawl basketball
	for {
		events := bc.Crawl()
		doc := bson.M{"date": time.Now().Format("2006-01-02"), "events": events}

		// insert data
		if _, err := collection.InsertOne(context.TODO(), doc); err != nil {
			zap.S().Error("failed to insert document:", err)
		} else {
			zap.S().Info("inserted a single document")
		}

		// wait for 10 second before next crawl
		time.Sleep(10 * time.Second)
	}

}
