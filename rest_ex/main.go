package main
import ("net/http" /* Paquetes para api rest o webservice */
        "log"
        )
func main(){
  /* Para el mux
   * Instalar gorilla/mux
   *  go get -u github.com/gorilla/mux
   */
  router:=NewRoute()
  /* Crea servidor forma 1
  http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"Hola mundo desde server go")
  })*/
  /* Escucha y sirve */
  server:=http.ListenAndServe(":8080",router)
  /* Show error */
  log.Fatal(server)
}
