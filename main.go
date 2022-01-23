package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/devaliakbar/gin_gorm_example/internal/core/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	server.RunServer()
}
