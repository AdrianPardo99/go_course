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
    Username: "d3vcr4ck",
    Password: "adrianPardo_99",
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
	m1:=Movie{"La insoportable levedad del ser",1988,"Philip Kaufman"}
	m2:=Movie{"John Wick",2014,"Chad Stahelski"}
	m3:=Movie{"John Wick 2",2017,"Chad Stahelski"}
	m4:=Movie{"John Wick Parabellum",2019,"Chad Stahelski"}
	movies:=[]interface{}{m1,m2,m3,m4}
	insert_result,err:=collection.InsertMany(context.TODO(),movies)
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
