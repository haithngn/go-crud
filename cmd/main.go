package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haithngn/go-crud/configuration"
	"github.com/haithngn/go-crud/db"
	"github.com/haithngn/go-crud/router"
	"github.com/haithngn/go-crud/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	fmt.Println("Application starting...")

	fmt.Println("Loading the env")
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("The env cannot be loaded %v", err)
	}

	config, err := configuration.GetConfig()
	if err != nil {
		log.Fatal("env cannot be loaded")
	}

	log.Println("Setting up database")
	database, err := db.Start(config.DBConfig())
	if err != nil {
		log.Fatalf("The database cannot be establish %v", err)
	}

	seeder := db.Seeder{}
	seeder.Populate(database)

	log.Println("Serving API...")

	engine := gin.Default()
	route := server.Server{Router: engine}

	route.Serve(router.GetRoutes(database))

}
