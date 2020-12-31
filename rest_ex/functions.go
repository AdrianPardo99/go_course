package main
import ("fmt"
        "net/http" /* Paquetes para api rest o webservice */
        "github.com/gorilla/mux"
        "encoding/json"
        "log"
        )
var movies=Movies{
  Movie{"La insoportable levedad del ser",1988,"Philip Kaufman"},
  Movie{"John Wick",2014,"Chad Stahelski"},
  Movie{"John Wick 2",2017,"Chad Stahelski"},
  Movie{"John Wick Parabellum",2019,"Chad Stahelski"},
}
func Index(w http.ResponseWriter,r *http.Request){
  fmt.Fprintf(w,"Bienvenido a un servidor hecho con Go")
}

func MovieList(w http.ResponseWriter,r *http.Request){

  json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter,r *http.Request){
  params:=mux.Vars(r)
  movie_id:=params["id"]
  fmt.Fprintf(w,"Has cargado la pelicula #%s",movie_id)
}

func MovieAdd(w http.ResponseWriter,r *http.Request){
  decoder:=json.NewDecoder(r.Body)
  var movie_data Movie
  err:=decoder.Decode(&movie_data)
  if err!=nil{
    panic(err)
  }
  defer r.Body.Close()
  log.Println(movie_data)
  movies=append(movies,movie_data)
}
