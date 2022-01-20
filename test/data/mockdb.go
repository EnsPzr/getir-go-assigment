package data

import (
	"context"
	"fmt"
	"github.com/strikesecurity/strikememongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"testing"
	"time"
)

// InitMockDB
// Starts a fake mongo instance in-memory for database and api test.
// Collection creates.
// Inserts mock test data into created collection.
// If don't any error, return fake mongo url, database name and mongo client.
func InitMockDB(t *testing.T) (string, string, *strikememongo.Server) {
	// Start mongo instance.
	mongoServer, err := strikememongo.StartWithOptions(&strikememongo.Options{
		MongoVersion:   "4.0.5",
		StartupTimeout: time.Duration(10) * time.Minute,
	})
	if err != nil {
		t.Fatal("Mongo server create error => ", err.Error())
	}

	// Read random database name.
	mongoURI := mongoServer.URIWithRandomDB()
	splitedDatabaseName := strings.Split(mongoURI, "/")
	databaseName := splitedDatabaseName[len(splitedDatabaseName)-1]

	uri := fmt.Sprintf("%s%s", mongoURI, "?retryWrites=false")
	// Connect db
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal("Mongo client create error => ", err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		t.Fatal("Mongo client connect error => ", err.Error())
	}

	database := client.Database(databaseName)
	// Create collection
	err = database.CreateCollection(context.Background(), "records")
	if err != nil {
		t.Fatal("mongo create collection error =>", err.Error())
	}
	// Get mocd data
	records, err := getMockRecordData()
	if err != nil {
		t.Fatal("Get Mock Record Error => " + err.Error())
	}
	// Insert mock data in records collection.
	for _, mockRecord := range records {
		_, err = database.Collection("records").InsertOne(ctx, mockRecord)
		if err != nil {
			t.Fatal("mongo create record error =>", err.Error())
		}
	}

	// Find all data in records collection.
	cur, err := database.Collection("records").Find(ctx, bson.D{})
	if err != nil {
		t.Fatal("Find collection data error => ", err.Error())
	}
	err = cur.All(ctx, &records)
	if err != nil {
		t.Fatal("cursor All error => ", err.Error())
	}
	return uri, databaseName, mongoServer
}
