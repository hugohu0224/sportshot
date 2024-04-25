package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
)

func InitMongodb(uri string) *mongo.Client {

	// 设置 MongoDB 连接字符串

	clientOptions := options.Client().ApplyURI(uri)

	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
	//"
	//	fmt.Println("Connected to MongoDB!")
	//
	//	// 获取数据库和集合
	//	collection := client.Database(databaseName).Collection(collectionName)
	//
	//	// 创建一个用户文档
	//	user := bson.D{
	//		{Key: "name", Value: "John Doe"},
	//		{Key: "age", Value: 30},
	//		{Key: "email", Value: "johndoe@example.com"},
	//	}
	//
	//	// 插入文档到集合中
	//	insertResult, err := collection.InsertOne(context.TODO(), user)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	//
	//	// 断开连接
	//	err = client.Disconnect(context.TODO())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println("Connection to MongoDB closed.")"
}
