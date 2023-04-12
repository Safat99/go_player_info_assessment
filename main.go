package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"player_info/routes"
	"player_info/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Welcome to the System")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// initialize the gin router
	router := gin.Default()

	// initializing the mongo client
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// db := client.Database("player_info")

	//mvc components
	playerService := &service.PlayerService{}

	//define routes
	routes.SetupPlayerRoutes(router, playerService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

}
