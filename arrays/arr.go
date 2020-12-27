package main
import "fmt"
func main(){
  var arr_str [4]string
  arr_str[0]="Parte 1"
  arr_str[1]="Parte 2"
  arr_str[2]="Parte 3"
  arr_str[3]="Parte 4"
  /* Imprime todos los elementos */
  fmt.Println(arr_str)

  /* Forma 2 */
  peliculas:=[3]string{"Star Wars","Inglourious Basterds","X Men"}
  fmt.Println(peliculas)
  /* Para imprimir solo 1 elemento */
  fmt.Println(peliculas[1])
}
