// Package main Beer API.
//
// The purpose of this application is to provide service of the value of a beer and the conversion of beers by type of currency
//
//
// This should show the struct of endpoints
// Terms Of Service:
//
//     Schemes: http, https
//     Host: localhost:8888
//     BasePath: /api/v1
//     Version: 1.0.0
//     Contact: https://www.linkedin.com/in/daniel-de-la-pava-suarez/
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/samuskitchen/beer-api-clean-arch/infrastructure/router"
)

func main() {
	log.Println("stating API cmd")
	port := os.Getenv("API_PORT")
	router.Start(port)
}
