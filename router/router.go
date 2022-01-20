package router

import (
	"github.com/enspzr/getir-go-assigment/handlers"
	"log"
	"net/http"
)

// Setup
// This function defines routes.
func Setup() {
	log.Println("Router setup started")

	http.HandleFunc("/records", handlers.RecordGetAll)
	http.HandleFunc("/in-memory", handlers.InMemory)
	http.HandleFunc("/in-memory-sqlite", handlers.InMemorySqlite)

	log.Println("Router setup finished")
}
