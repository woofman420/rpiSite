package main

import (
	"rpiSite/database"
	"rpiSite/handlers"
)

func main() {
	database.Initialize()
	handlers.Initalize()
}
