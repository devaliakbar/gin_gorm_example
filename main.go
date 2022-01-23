package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/devaliakbar/gin_gorm_example/internal/core/server"
)

func main() {
	server.RunServer()
}
