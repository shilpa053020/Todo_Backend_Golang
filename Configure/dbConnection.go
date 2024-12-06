package configure

import (
	utils "Todo_Backend_Golang/Utils"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB client
var Client *mongo.Client

// ConnectToDb initializes the connection to MongoDB
func ConnectToDb() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve DBURL from environment
	DbUrl := os.Getenv("DBURL")
	if DbUrl == "" {
		log.Fatal("DBURL is not set in the environment")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(DbUrl)
	ctx, cancel := utils.GetContext(10*time.Second)
	defer cancel()

	// Connect to MongoDB
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect with DB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")
}
