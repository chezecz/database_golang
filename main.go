package track_mongo 

import (
	"../uts_course"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	"github.com/fatih/structs"
)

func Search_subject(course course_retrieve.Course) {

}

func Insert_course(course course_retrieve.Course) {

	fmt.Println("I am INSERTING")
	var mongo_url = "mongodb://localhost:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url));

	if err != nil {
		log.Fatal(err);
	}

	fmt.Println("Connected to the mongo database ", mongo_url)

	courses := client.Database("track").Collection("courses");

	mapped_course := structs.Map(course)

	res, err := courses.InsertOne(context.TODO(), bson.M(mapped_course))

	fmt.Println(bson.M(mapped_course))

	fmt.Println(res, "\n", err)

	err = client.Disconnect(context.TODO());
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println("\nDISCONNECTED");
}

func Insert_subject(subject course_retrieve.Subject) {

}

func Insert_requirements(requirements []course_retrieve.RequirementGroup) {

}

func Get_requirements (course course_retrieve.Course) []course_retrieve.RequirementGroup {
	return nil
}

func Search_course (course_id string) course_retrieve.Course {

	fmt.Println("I AM HERE")

	var mongo_url = "mongodb://localhost:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url));

	if err != nil {
		log.Fatal(err);
	}

	fmt.Println("Connected to the mongo database ", mongo_url)

	courses := client.Database("track").Collection("courses");


	//subjects = client.Database("track").Collection("subjects");

	filter := bson.D{{"CourseId", course_id}};

	var search course_retrieve.Course;
	err = courses.FindOne(context.TODO(), filter).Decode(&search);	

	fmt.Println(search)

	err = client.Disconnect(context.TODO());
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println("\nDISCONNECTED");

	return search



}
