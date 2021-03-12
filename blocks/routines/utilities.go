package utilities

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnetMongo() {
	localConnectStr := "mongodb://localhost:27017/?readPreference=primary&ssl=false"
	client, err := mongo.NewClient(options.Client().ApplyURI(
		// remoteConnectStr,
		localConnectStr,
	))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	db := client.Database("local")
	collection := db.Collection("startup_log")
	cursor, err := collection.Find(ctx, bson.M{"_id": "ubuntu-1615493216307"})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var row bson.M
		if err = cursor.Decode(&row); err != nil {
			log.Fatal(err)
		}
		fmt.Println(row["hostname"], " : ", row["cmdLine"])
		fmt.Println("---------------------")
		break
	}
}

func Double(a int) int {
	return a * 2
}
