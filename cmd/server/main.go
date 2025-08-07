package main

import (
	"Runlet/internal/infrastructure"
	"Runlet/internal/interfaces/api"
)

func main() {
	dbClient := infrastructure.GetDbClient()
	router := api.GetRouter(dbClient)
	router.Run(":8080")
}
