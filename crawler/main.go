package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"log"
	"os"
	"sportshot/db/mongodb"
	"time"
)

func main() {
	// initialize logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	// initial basketball crawler
	bc := basketballCrawler{}
	url := "https://tw.betsapi.com/ciz/basketball"
	events := bc.crawl(url)

	// print to make sure we get the data for dev test
	jsonData, err := json.MarshalIndent(events, "", "    ")
	if err != nil {
		log.Println("Error marshaling data:", err)
		return
	}
	fmt.Println("Data extracted:\n", string(jsonData))

	// get mongo config
	data, err := os.ReadFile("./crawler/config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	uri := config["mongodbURI"].(string)
	client := mongodb.InitMongodb(uri)

	defer func(mongoc *mongo.Client, ctx context.Context) {
		err := mongoc.Disconnect(ctx)
		if err != nil {
			zap.S().Fatal("Error disconnecting from mongodb:", err)
		}
	}(client, context.TODO())
	databaseName := "sportevents"
	collectionName := "basketball"

	collection := client.Database(databaseName).Collection(collectionName)
	doc := bson.M{"date": time.Now().Format("2001-01-01"), "events": events}
	insertResult, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}
