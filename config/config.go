package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MONGO_DRIVER_DOC = https://www.mongodb.com/docs/drivers/go/current/quick-start/

var db_Name string = "golang-core"

func DatabaseConfig() *mongo.Client {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error:", err)
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal("Error:", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

var Client *mongo.Client = DatabaseConfig()

// following code to create a database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(db_Name).Collection(collectionName)
	return collection
}
