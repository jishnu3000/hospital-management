package config

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Read MongoDB URI from environment variable
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatal("MONGODB_URI not set in .env file")
    }

    // Connect to MongoDB
    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    // Check if the connection is successful
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal("Database connection failed!")
    }

    log.Println("Connected to MongoDB!")
    DB = client.Database("hospital_management")
}
