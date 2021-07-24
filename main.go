package main

import (
	"log"

	"github.com/cbuelvasc/twittor/database"
	"github.com/cbuelvasc/twittor/handlers"
)

func main() {
	if database.CheckConnection() == 0 {
		log.Fatal("No database connection")
		return
	}
	handlers.Handlers()
}
