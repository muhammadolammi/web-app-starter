package main

import "web-app-starter/internal/database"

type Config struct {
	PORT string
	DB   *database.Queries
}
