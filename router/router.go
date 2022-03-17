package router

import (
	_ "github.com/enspzr/getir-go-assigment/docs"
	"github.com/enspzr/getir-go-assigment/handlers"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// Setup
// This function defines routes.
func Setup() {
	log.Println("Router setup started")

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/records", handlers.RecordGetAll)
	http.HandleFunc("/in-memory", handlers.InMemory)
	http.HandleFunc("/in-memory-sqlite", handlers.InMemorySqlite)

	log.Println("Router setup finished")
}
