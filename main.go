package main

import (
	"database/sql"
	"log"
	"os"
	"web-app-starter/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("This is the starter pack for golang servers built for website backend")
	//   Lets load the environment
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("there is no port provided kindly provide a port.")
		return
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Println("empty dbURL")
		return
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println(err)
		return
	}
	dbQueries := database.New(db)

	config := Config{
		PORT: port,
		DB:   dbQueries,
	}

	config.server()

}
