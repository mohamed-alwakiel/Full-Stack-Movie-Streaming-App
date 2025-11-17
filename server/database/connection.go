package database

import (
	"fmt"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBConnect() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		logger.Fatal("MONGODB_URI not set !!")
	}
	fmt.Println("Connected to MongoDB! : " + mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		logger.Info("failed to connect to MongoDB")
		return nil
	}
	return client
}

var Client *mongo.Client = DBConnect()

func OpenConnection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file")
	}
	DBName := os.Getenv("DATABASE_NAME")
	fmt.Println("DATABASE_NAME : " + DBName)

	collection := Client.Database(DBName).Collection(collectionName)
	if collection == nil {
		logger.Error("Failed to connect to MongoDB with collection " + collectionName)
		return nil
	}
	return collection
}
