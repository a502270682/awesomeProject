package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func getConn() *mongo.Client{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return client
}

func StartTran() {
	ctx := context.Background()
	client := getConn()
	coll := client.Database("test").Collection("test")
	var err error
	err = client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction()
		if err != nil {
			fmt.Println(err)
			return err
		}

		//在事务内写一条id为“222”的记录
		_, err = coll.InsertOne(sessionContext, bson.M{"_id": "222", "name": "ddd", "age": 50})
		if err != nil{
			fmt.Println(err)
			return err
		}

		//在事务内写一条id为“333”的记录
		_, err = coll.InsertOne(sessionContext, bson.M{"_id": "333", "name": "ddd", "age": 50})
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			return err
		} else {
			sessionContext.CommitTransaction(sessionContext)
		}
		return nil
	})
}
