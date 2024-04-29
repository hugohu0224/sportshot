package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"log"
	"os"
	"sportshot/crawler/global"
	"sportshot/crawler/initial"
	"sportshot/crawler/operator"
	"time"
)

func main() {
	// initial logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infof("Logger initialized")

	// initial basketball crawler
	bc := operator.BasketballCrawler{}
	zap.S().Infof("BasketballCrawler initialized")

	// read config
	data, err := os.ReadFile("config.json")
	if err != nil {
		zap.S().Fatalf("Error reading config.json: %v", err)
	}
	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	zap.S().Infof("Loaded config %v", config)

	// initial MongoClient
	uri := config["mongodbURI"].(string)
	initial.InitMongoClient(uri)
	defer func(mongoc *mongo.Client, ctx context.Context) {
		err := mongoc.Disconnect(ctx)
		if err != nil {
			zap.S().Fatal("Error disconnecting from mongodb:", err)
		}
	}(global.MongodbClient, context.TODO())
	zap.S().Infof("MongoClient initialized")

	// connect to mongo
	databaseName := "sportevents"
	collectionName := "basketball"
	collection := global.MongodbClient.Database(databaseName).Collection(collectionName)
	zap.S().Infof("Mongodb.%s.%s connected", databaseName, collectionName)

	// start to crawl basketball
	for {
		url := "https://tw.betsapi.com/ciz/basketball"
		events := bc.Crawl(url)
		doc := bson.M{"date": time.Now().Format("2006-01-02"), "events": events}

		// insert data
		if _, err := collection.InsertOne(context.TODO(), doc); err != nil {
			zap.S().Error("Failed to insert document:", err)
		} else {
			zap.S().Info("Inserted a single document")
		}

		// wait for 10 second before next crawl
		time.Sleep(10 * time.Second)
	}

}
