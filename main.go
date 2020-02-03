package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ashishprasad2163:05222412532@cluster0-fzcwc.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	DockDatabase := client.Database("Dock")
	podcastCollection := DockDatabase.Collection("podcast")
	episodeCollection := DockDatabase.Collection("episode")
	/*podcastResult, err := podcastCollection.InsertOne(ctx, bson.D{
		{Key: "Title", Value: "Ashiqui"},
		{Key: "Author", Value: "Arijit"},
		{"tags", bson.A{"singer", "vocal", "artists"}},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcastResult.InsertedID)
	episodeResult, err := episodeCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"Title", "Volume1"},
			{"description", "Volume number 1"},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"Title", "Volume2"},
			{"description", "Volume number 2"},
			{"duration", 35},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodeResult.InsertedIDs)*/

	/*//Retrieve and query all in batches

	cursor, err := episodeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episode bson.M
		err = cursor.Decode(&episode)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(episode)
	}*/

	//retreive single document
	var podcast bson.M
	if err = podcastCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcast)

	//reteive using filter
	var episodesFiltered []bson.M
	filtercursor, err := episodeCollection.Find(ctx, bson.M{"duration": 35})
	if err = filtercursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)

	ops := options.Find()
	ops.SetSort(bson.D{{"duration", -1}})
	sortCursor, err := episodeCollection.Find(ctx, bson.D{
		{"duration", bson.D{
			{"$gt", 24},
		}},
	}, ops)

	var episodesSorted []bson.M
	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesSorted)

	/*//update db with set Id of one document

	id, _ := primitive.ObjectIDFromHex("5e356c4b4ab88b2075d3de6c")
	result, err := podcastCollection.UpdateOne(
		ctx, bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"Author", "Neha Kakkar"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated the docs", result.ModifiedCount)*/

	/*//update db with title of one and multi documents
	result, err := podcastCollection.UpdateMany(
		ctx, bson.M{"Title": "Ashiqui"},
		bson.D{
			{"$set", bson.D{{"Author", "Neha Kakkar and Guru"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated the docs", result.ModifiedCount)*/

	//full on replace of docs
	/*result, err := podcastCollection.ReplaceOne(
		ctx, bson.M{"Title": "Ashiqui"},
		bson.M{
			"Title":  "Ashiqui 2",
			"Author": "Chetan BHagat"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Replaced the docs", result.ModifiedCount)*/

	//delete
	/*result, err := podcastCollection.DeleteOne(
		ctx, bson.M{"Title": "Ashiqui 2"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted the docs", result.DeletedCount)*/

}
