package main
import "fmt"
func main(){
  series:=[]string{"Dr_House","The_Office"}
  fmt.Println(series)
  fmt.Println("Tamaño del arreglo: ",len(series))
  /* Para añadir un nuevo elemento tendremos */
  series=append(series,"Dexter")
  fmt.Println(series)
  fmt.Println("Tamaño del arreglo: ",len(series))
}
