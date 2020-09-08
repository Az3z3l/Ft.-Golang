package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

	//  "_id" : ObjectId("5f51c80128be046a772c963c"),
	// "id" : 0,
	// "username" : "uno",
	// "password" : "asta",
	// "email" : "Assd@asd.cd",
	// "key" : "YlliVRoYbVvpXoADYbixSeiYo29z77.5rWcvoToXSvHtbfYDqWVyKUG7ikcD3Lkt2JrMk42K7zuUFM1pv5MG5hZ0g00xhAq83b4WjfjX1tgV6IAbSbKvkparYwGRNKNJ",
	// "solved" : null,
	// "isadmin" : true

	type User struct {
		id       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Key      string `json:"key"`
		solved   string `json:"solved"`
		Isadmin  *bool  `json:"isadmin"`
	}

	collection := client.Database("ctf").Collection("players")
	filter := bson.D{{"username", "username"}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())
	var result User
	for cursor.Next(context.TODO()) {
		var v User
		err := cursor.Decode(&v)
		if err != nil {
			panic(err)
		}
		result = v
	}
	fmt.Print(result)

	var v User
	err = collection.FindOne(context.TODO(), filter).Decode(&v)
	fmt.Println(v)

}
