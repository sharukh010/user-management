package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetSession() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occurred while loading env")
	}

	connString := os.Getenv("MONGODB_URL")
	fmt.Println("Connecting to MongoDB at:", connString)

	clientOpts := options.Client().ApplyURI(connString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOpts)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	fmt.Println("✅ MongoDB session created successfully")
	return client
}
