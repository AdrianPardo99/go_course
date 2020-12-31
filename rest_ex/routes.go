package main
import ("net/http" /* Paquetes para api rest o webservice */
        "github.com/gorilla/mux"
)
type Route struct{
  Name string
  Method string
  Path string
  HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRoute() *mux.Router{
  router:=mux.NewRouter().StrictSlash(true)
  for _, route:=range routes{
    router.
            Name(route.Name).
            Methods(route.Method).
            Path(route.Path).
            Handler(route.HandleFunc)
  }
  return router
}

var routes=Routes{
  Route{"Index","GET","/",Index,},
  Route{"MovieList","GET","/peliculas",MovieList,},
  Route{"MovieShow","GET","/pelicula/{id}",MovieShow,},
  Route{"MovieAdd","POST","/pelicula",MovieAdd,},
}
