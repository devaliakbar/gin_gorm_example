package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/devaliakbar/gin_gorm_example/internals/core/server"
)

func main() {
	server.RunServer()
}
