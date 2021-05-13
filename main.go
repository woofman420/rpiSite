package main

import (
	"log"
	"rpiSite/database"
	"rpiSite/handlers"
)

func main() {
	if err := database.Initialize(); err != nil {
		log.Fatal(err)
	}
	handlers.Initialize()
}
