package main

import (
	"context"
	"fmt"
	"log"
	"player_info/config"
	"player_info/repository"
	"player_info/routes"
	"player_info/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Welcome to the System")

	// loading configurations
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(cfg.DBURI)
	//create client using those configs
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.DBURI))
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
	}

	// now connect that client
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer client.Disconnect(context.Background())

	playerRepository := repository.NewPlayerRepository(client, cfg.DBName)
	playerService := service.NewPlayerService(playerRepository)

	//set up router and start listening
	// routes.SetupPlayerRoutes(router, playerService)
	r := routes.SetupPlayerRouter(playerService)
	r.Run(":8080")

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	// log.Printf("Server listening on port %s", port)
	// err = http.ListenAndServe(":"+port, router)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
