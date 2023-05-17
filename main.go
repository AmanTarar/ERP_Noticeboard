package main

import (
	"fmt"
	"log"
	"main/server"
	"main/server/db"

	// "main/server/socket"
	"os"

	"github.com/joho/godotenv"
)

// @title Gin Demo App
// @version 1.0
// @description This is a demo version of Gin app.
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		
	}
	fmt.Println("env var loaded")

	 MongoCollection := db.ConnectDB()
	db.Collection=MongoCollection

	

	app := server.NewServer()
	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}





  
 