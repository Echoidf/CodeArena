package mongodb

import (
	"content-assistant/pkg/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func init() {
	// 初始化数据库
	uri := config.GetConfig().MongoDBUri
	DB = InitDb(uri).Database("content-assistant")
}

func InitDb(uri string) *mongo.Client {
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	return client
}
