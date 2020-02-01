package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// To Do:
// establish a connection to mongodb
// define our data model
// define endpoints

type Cat struct {
	// Mongodb's object id
	ID    primitive.ObjectID `json:_id,omitempty bson:_id,omitempty`
	Name  string             `json:name,omitempty bson:name,omitempty`
	Color string             `json:color,omitempty bson:color,omitempty`
}

// create an instance for the client connection
var client *mongo.Client

func main() {
	fmt.Println("yay")
	// establish connection, with timeouts
	// context (here skipping errors)
	client, _ = mongo.NewClient(
		options.Client().ApplyURI("mongodb://localhost").SetConnectTimeout(10 * time.Second))
	// define router (mux, not the standard go http)
	router := mux.NewRouter()
	http.ListenAndServe(":12345", router)
}
