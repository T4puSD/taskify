package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient = connectDB()

func connectDB() *mongo.Client {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error occured during loading env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetMaxPoolSize(10))

	if err != nil {
		log.Fatal("Unable to build mongodb client", err)
	}

	ctx, cancel := GetContext()
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal("Unable to connect to mongodb: ", err)
	}

	log.Println("Connected to mongodb!!!")
	return client
}

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Millisecond*10)
}

func GetCollection(collectionName string) *mongo.Collection {
	return dbClient.Database(os.Getenv("MONGODB_DATABASE_NAME")).Collection(collectionName)
}
