package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"sportshot/crawler/operator"
	"sportshot/utils/config"
	"sportshot/utils/db"
	"sportshot/utils/global"
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
		time.Sleep(10 * time.Second)
	}
}
