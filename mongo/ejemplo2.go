package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

/* Modelo */
type Movie struct{
  Name string     `json:"name"`
  Year int        `json:"year"`
  Director string `json:"director"`
}
/* Objeto generalizado / Slice */
type Movies []Movie

func main() {
	host := "localhost"
	port := 27017
	credential := options.Credential{
    Username: "username",
    Password: "password",
	}
	clientOpts := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)).
		SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	/* Check the connections */
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
  collection:=client.Database("curso_go").Collection("movies")

  /* Aquí realizaremos la consulta de peliculas */
  filtro:=bson.D{{"name","John Wick 2"}}
  updateResult, err := collection.UpdateOne(context.TODO(), filtro, bson.D{
    {"$set", bson.D{
      {"name", "John Wick Chapter 2"},
      }},
  })
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Matched %v documents and updated %v documents.\n",
    updateResult.MatchedCount, updateResult.ModifiedCount)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
fmt.Println("Closed connection")

	fmt.Println("Congratulations, you're already connected to MongoDB!")
}
