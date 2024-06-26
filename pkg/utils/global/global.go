package global

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	MongodbClient *mongo.Client
	EtcdClient    *clientv3.Client
	MySQLClient   *gorm.DB
)
