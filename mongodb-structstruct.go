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

	type Userwhen struct {
		ChallID   string `json:"ChallID"`
		Timestamp string `json:"Timestamp"`
	}

	type Challwhen struct {
		UserID    string `json:"UserID"`
		TimeStamp string `json:"TimeStamp"`
	}

	type User struct {
		ID       string      `bson:"_id"`
		Username string      `json:"username"`
		Password string      `json:"password"`
		Email    string      `json:"email"`
		Solved   []*Userwhen `json:"solved"`
		Score    int         `json:"score"`
	}

	type Challenge struct {
		ID          string       `bson:"_id"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Category    []*string    `json:"category"`
		Value       int          `json:"value"`
		Flags       string       `json:"flags"`
		Solves      *int         `json:"solves"`
		Teamssolved []*Challwhen `json:"teamssolved"`
	}

	Chacollection := client.Database("ctf").Collection("challenges")
	Placollection := client.Database("ctf").Collection("players")
	id := "5f5bb5e18d9e7a44959d167a"
	docid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic("Cannot find object ID")
	}
	var chall *Challenge
	var playa *User
	flag := "flag{100-pt-chall}"
	filter := bson.M{"flags": bson.M{"$eq": flag}}
	err = Chacollection.FindOne(nil, filter).Decode(&chall)
	if err != nil {
		panic("Please check your Flag again")
	}
	// chall HAS NOW THE CHALLENGE THAT HAS THE FLAG
	challid, err := primitive.ObjectIDFromHex(chall.ID)
	fmt.Println(chall)
	filter = bson.M{"_id": bson.M{"$eq": docid}}
	err = Placollection.FindOne(nil, filter).Decode(&playa)
	if err != nil {
		panic(err)
	}
	// this is the first solve for the team ..... create an array inside the key --> solved
	if playa.Solved == nil {
		update := bson.M{"$set": bson.M{"solved": []interface{}{
			bson.M{
				"ChallID":   challid,
				"Timestamp": "12/12/12",
			}}}}
		_, erro := Placollection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if erro != nil {
			panic(erro)
		}
	} else { // already solved atleast one challenge ..... see if this challenge is solved already
		for _, i := range playa.Solved {
			if i.ChallID == chall.ID {
				fmt.Println("Already solved this Challenge") // use return here
				panic("yolo")
			}
		}
	}
	// this challenge not solved yet
	update := bson.M{"$push": bson.M{"solved": bson.M{"ChallID": challid, "Timestamp": "12/20/2020"}}}
	_, erro := Placollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		panic(erro)
	}
	// Added data to USERS side =>  Make the changes in challenges side
	filter = bson.M{"_id": bson.M{"$eq": challid}}
	// First blood needs spl attention ~ create an array before pushing
	if chall.Teamssolved == nil {
		update := bson.M{"$set": bson.M{"Teamssolved": []interface{}{
			bson.M{
				"UserID":    docid,
				"Timestamp": "12/12/12",
			}}}}
		_, erro = Chacollection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if erro != nil {
			panic(erro)
		}
	} else {
		update = bson.M{"$push": bson.M{"Teamssolved": bson.M{"UserID": docid, "Timestamp": "12/20/2020"}}}
		_, erro = Chacollection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if erro != nil {
			panic(erro)
		}
	}
	// Increase score for playa
	filter = bson.M{"_id": bson.M{"$eq": docid}}
	score := playa.Score + chall.Value
	update = bson.M{"$set": bson.M{"Score": score}}
	_, erro = Placollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if erro != nil {
		fmt.Println("Error updating")
	}
	fmt.Println("Correct Flag")

}

// db.players.update({"_id":ObjectId("5f59cc3b2a9ec7c8ec10068d")}, {$push : {"solved":{"ChallID":ObjectId("5f5780da23104197b1306ba5"), "Timestamp":"12/12/2020"}} })

// db.players.update({"_id":ObjectId("5f59cc3b2a9ec7c8ec10068d")}, {$set : {"solved" : [{"ChallID":ObjectId("5f5780da23104197b1306ba5"), "Timestamp":"12/12/2020"}]} })
