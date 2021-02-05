package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson"
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

	/* Vereficamos al driver hacia donde haremos la alta... */
	collection:=client.Database("curso_go").Collection("movies")

	/* Insercion de datos de ejemplo */
	var movies=Movies{
    Movie{"La insoportable levedad del ser",1988,"Philip Kaufman"},
	  Movie{"John Wick",2014,"Chad Stahelski"},
	  Movie{"John Wick 2",2017,"Chad Stahelski"},
	  Movie{"John Wick Parabellum",2019,"Chad Stahelski"},
  }
	movies_s:=[]interface{}{}
  for i:=0;i<len(movies);i++{
    movies_s=append(movies_s,movies[i])
  }
  fmt.Println(movies_s)
	insert_result,err:=collection.InsertMany(context.TODO(),movies_s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movies had been inserted: ", insert_result.InsertedIDs)
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
fmt.Println("Closed connection")

	fmt.Println("Congratulations, you're already connected to MongoDB!")
}
