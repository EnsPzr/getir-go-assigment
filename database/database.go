package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"time"
)

var db *mongo.Client
var dbName string

func DB() *mongo.Client {
	return db
}

func DbName() string {
	return dbName
}

// Connect
// This function opens database connection.
// If any error, return error.
func Connect(uri string) error {
	log.Println("Database connect started")
	var err error

	db, err = mongo.NewClient(options.Client().
		ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Connect(ctx)
	if err != nil {
		return err
	}
	err = setDbName(uri)
	if err != nil {
		return err
	}

	log.Println("Database connect finished")
	return nil
}

// This function read database name in connection url.
// And sets database name in global variable.
func setDbName(uri string) error {
	splitedDatabaseName := strings.Split(uri, "/")
	if len(splitedDatabaseName) == 0 {
		return errors.New("Db Name not found")
	}
	dName := strings.Split(splitedDatabaseName[len(splitedDatabaseName)-1], "?")
	if len(dName) == 0 {
		return errors.New("Db Name not found")
	}
	dbName = dName[0]
	return nil
}

// Disconnect
// This function closes database connection.
func Disconnect() {
	db.Disconnect(context.Background())
	log.Println("Database connection closed")
}
