package main

import (
	"fmt"
	"github.com/enspzr/getir-go-assigment/cache"
	"github.com/enspzr/getir-go-assigment/database"
	"github.com/enspzr/getir-go-assigment/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// This is start function.
// Many operations are performed in this function.
// 1- Database connection is opened. If doesn't connect database, project stops.
// 2- Go-cache initializes.
// 3- Sqlite cache initilizes. If doesn't initialize Sqlite cache, project stops.
// 4- Router are defined.
// 5- Http server start and listens request.
func main() {
	log.Println("Backend Starting Up")
	err := database.Connect("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true")
	if err != nil {
		log.Fatalf("Database init error: " + err.Error())
	}
	defer database.Disconnect()

	cache.InitGoCache()

	err = cache.InitSqliteCache()
	if err != nil {
		log.Fatalf("Sqlite cache init error: " + err.Error())
	}

	router.Setup()

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080" // Default port if not specified
		}
		err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Backend Started Succesfully")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	log.Print("Gracefully shutting down")
}
