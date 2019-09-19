package main

import (
	"fmt"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"index.go"
)
type StudentPassedSubject struct {
	Name string
	Subject_id string
	Grade string
}
type StudentPlannedSubject struct {
	Name string
	Subject_id string
}
type StudentSubjects struct {
	Passed []StudentPassedSubject
	Planned []StudentPlannedSubject
}

type StudentCourse struct {
	Name string
	Major string
	Id string
}

type Student struct {
	Name string
	Email string
	Student_id string
	Course StudentCourse
	Subject StudentSubjects
}

/*
type Colors struct {
	Color string
}
*/

func main () {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"));

	fmt.Println("I AM CONNECTING\n");

	if err != nil {
		log.Fatal(err);
	}

	err = client.Ping(context.TODO(), nil);
	if err != nil {
		log.Fatal(err);
	}

	fmt.Println("CONNECTED\n");

	collection := client.Database("test").Collection("students");


/*
//HERE COMES INSERT CODE
	newColor := Colors{"pink"}
	insertResult, err := collection.InsertOne(context.TODO(), newColor);
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println("Inserted: ", insertResult.InsertedID);

*/


	filter := bson.D{{"name", "Aleksandr Gromov"}};

	//var search Colors;
	var search Student;

	err = collection.FindOne(context.TODO(), filter).Decode(&search);

	fmt.Printf("Found a single document: %+v\n", search);


	filter = bson.D{{"name", "Master of Information Technology"}};

	var courseSearch 


	err = client.Disconnect(context.TODO());
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println("\nDISCONNECTED");
}
