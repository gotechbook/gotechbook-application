package mongo

import (
	"context"
	"github.com/gotechbook/gotechbook-application/config"
	"github.com/topfreegames/pitaya/v2/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONFIG = "config_db"
	CHAT   = "chat_db"
)

func GetMongoClient(ctx context.Context) (c *mongo.Client) {
	if err := config.GOTECHBOOK_MONGO.Ping(ctx, nil); err != nil {
		logger.Log.Fatal("mongodb-err")
	}
	return config.GOTECHBOOK_MONGO
}
func Insert(ctx context.Context, client *mongo.Client, db string, tb string, value interface{}) (insertResult *mongo.InsertOneResult, err error) {
	collection := client.Database(db).Collection(tb)
	return collection.InsertOne(ctx, value)
}
func Update(ctx context.Context, client *mongo.Client, db string, tb string, filter interface{}, update interface{}) (updateResult *mongo.UpdateResult, err error) {
	collection := client.Database(db).Collection(tb)
	return collection.UpdateMany(ctx, filter, bson.M{"$set": update})
}
func FindOne(ctx context.Context, client *mongo.Client, db string, tb string, filter interface{}, result interface{}, opts ...*options.FindOneOptions) (err error) {
	collection := client.Database(db).Collection(tb)
	return collection.FindOne(ctx, filter, opts...).Decode(result)
}
func FindCount(ctx context.Context, client *mongo.Client, db string, tb string, filter interface{}) (count int64, err error) {
	collection := client.Database(db).Collection(tb)
	count, err = collection.CountDocuments(ctx, filter)
	return
}
func FindAll(ctx context.Context, client *mongo.Client, db string, tb string, filter interface{}, result interface{}, opts ...*options.FindOptions) (rst interface{}, err error) {
	collection := client.Database(db).Collection(tb)
	cursor, err := collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
