package main

import (
	_ "github.com/BlackBird125/GoCRUD2/docs"

	"github.com/BlackBird125/GoCRUD2/db"
	"github.com/BlackBird125/GoCRUD2/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// @title gin-swagger todos
// @version 1.0
// @license.name kosuke
// @description このswaggerはgin-swaggerの見本apiです
func main() {
	loadEnv()
    db.Init()
    server.Init()
    db.Close()
}


func loadEnv() {
	godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
}