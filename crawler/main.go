package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"sportshot/crawler/global"
	"sportshot/crawler/operator"
	"sportshot/utils/config"
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

	// initial config
	config.InitConfigByViper()
	uri := db.GetMongodbURI()
	global.MongodbClient = db.GetMongodbClient(uri)
	defer func(c *mongo.Client, ctx context.Context) {
		err := c.Disconnect(ctx)
		if err != nil {
			zap.S().Fatal("error disconnecting from mongodb:", err)
		}
	}(global.MongodbClient, context.TODO())
	zap.S().Infof("mongoClient initialized")

	// start to crawl basketball odds
	for {
		events := bc.Crawl()
		bc.SaveToMongo(events)
		// wait for 10 second before next crawl
		time.Sleep(10 * time.Second)
	}
}
