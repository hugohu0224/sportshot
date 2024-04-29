package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"log"
	"os"
	"sportshot/crawler/db"
	"sportshot/crawler/operator"
	"time"
)

func main() {
	// initialize logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	// initial basketball crawler
	bc := operator.BasketballCrawler{}
	url := "https://tw.betsapi.com/ciz/basketball"

	// read config
	data, err := os.ReadFile("./crawler/config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	// get mongo info
	uri := config["mongodbURI"].(string)
	client := db.GetMongoClient(uri)
	defer func(mongoc *mongo.Client, ctx context.Context) {
		err := mongoc.Disconnect(ctx)
		if err != nil {
			zap.S().Fatal("Error disconnecting from mongodb:", err)
		}
	}(client, context.TODO())
	// connect to mongo
	databaseName := "sportevents"
	collectionName := "basketball"
	collection := client.Database(databaseName).Collection(collectionName)

	// start to crawl
	for {
		events := bc.Crawl(url)
		doc := bson.M{"date": time.Now().Format("2006-01-02"), "events": events}

		// insert data
		if _, err := collection.InsertOne(context.TODO(), doc); err != nil {
			zap.S().Error("Failed to insert document:", err)
		} else {
			zap.S().Info("Inserted a single document")
		}

		// Wait for 10 second before next crawl
		time.Sleep(10 * time.Second)
	}

}
