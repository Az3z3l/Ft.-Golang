package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		ID       string `bson:"_id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	collection := client.Database("ctf").Collection("players")
	filter := bson.D{{"username", "start7ricks"}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())
	var result []User
	for cursor.Next(context.TODO()) {
		var v User
		err := cursor.Decode(&v)
		if err != nil {
			panic(err)
		}
		result = append(result, v)

	}

	// fmt.Println(result)
	var f User
	codes := collection.FindOne(context.TODO(), filter).Decode(&f)
	if codes != nil {
		fmt.Println(codes)
	}
	// fmt.Println(f.ID)

	var u User
	id := "5f59cc3b2a9ec7c8ec10068D"
	id = "5f59cc3b2a9ec7c8ee10068D"
	docid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic("Cannot find object ID")
	}
	filtera := bson.M{"_id": bson.M{"$eq": docid}}
	err1 := collection.FindOne(context.TODO(), filtera).Decode(&u)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(u)
}
