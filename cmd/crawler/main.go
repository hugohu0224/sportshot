package main

import (
	"context"
	"go.uber.org/zap"
	"sportshot/internal/crawler/operator"
	"sportshot/pkg/utils/db"
	"sportshot/pkg/utils/global"
	"time"
)

func main() {
	// logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infof("logger initialized")

	// basketball crawler
	basketballCrawler := operator.BasketballCrawler{}
	zap.S().Infof("BasketballCrawler initialized")

	// config
	db.InitConfigByViper()

	// mongodb
	global.MongodbClient = db.GetMongodbClient()
	defer global.MongodbClient.Disconnect(context.TODO())
	zap.S().Infof("MongoClient initialized")

	// crawl
	for {
		zap.S().Infof("start to crawl")
		events := basketballCrawler.Crawl()
		basketballCrawler.SaveToMongo(events)
		time.Sleep(5 * time.Second)
	}
}
