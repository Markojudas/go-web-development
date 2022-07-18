package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// database
var DB *mongo.Database

//collections
var Books *mongo.Collection

func init() {
	//get a mongo session
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	session, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://bond:moneypenny007@localhost/bookstore"))
	if err != nil {
		panic(err)
	}

	if err = session.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	DB = session.Database("bookstore")
	Books = DB.Collection("books")

	fmt.Println("You connected to your mongo database")
}
