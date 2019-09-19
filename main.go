package main

import (
	"fmt"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/bson"
)	

type User struct {
	User_id int `bson:"userid" json:"userid"`
	Name string `bson:"name" json:"name"`
}

func main() {
	client, ctx := connect_db()

	collection := get_collection(client)

	res := insert_users(User{1, "Dendi"}, collection, ctx)
	
	fmt.Println(res)

	close_connection(client)
}

func connect_db() (mongo.Client, context.Context) {
	var err error

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		fmt.Println(err)
	}

	return *client, ctx
}

func get_collection(client mongo.Client) (mongo.Collection) {
	collection := client.Database("test_project").Collection("Users")
	return *collection
}

func insert_users(user User, collection mongo.Collection, ctx context.Context) (mongo.InsertOneResult) {
	var err error
	res, err := collection.InsertOne(ctx, user)

	if err != nil {
		fmt.Println(err)
	}

	return *res
}

// func find_users() {
// 	var result User
// 	err = collection.FindOne(ctx, filter).Decode(&result)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return result
// }

func close_connection(client mongo.Client) {
	var err error
	err = client.Disconnect(context.TODO())

	if err != nil {
		fmt.Println(err)
	}
}