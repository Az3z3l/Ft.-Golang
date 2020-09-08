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

	var results []*User
	collection := client.Database("ctf").Collection("players")
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	fmt.Println(results[0].id)
	fmt.Printf("Found multiple documents (array of pointers): %v+\n", &results)

	// fmt.Println(client)
	// fmt.Println(collection)

}
