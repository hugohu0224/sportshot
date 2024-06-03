package main

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sportshot/internal/crawler"
	"sportshot/pkg/utils/db"
	"sportshot/pkg/utils/global"
	"sportshot/pkg/utils/tools"
)

func main() {
	// logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infof("logger initialized")

	// basketball crawler
	basketballCrawler := &crawler.BasketballCrawler{}
	zap.S().Infof("BasketballCrawler initialized")

	// config
	db.InitConfigByViper()
	zap.S().Infof("config initialized")

	// mongodb
	global.MongodbClient = db.GetMongodbClient()
	defer global.MongodbClient.Disconnect(context.TODO())
	zap.S().Infof("MongoClient initialized")

	// start to crawl
	go tools.CrawlerTicker(basketballCrawler, viper.GetInt("CRAWL_SECOND_INTERVAL"))
	select {}

}
