package main
import "fmt"
func main(){
  fmt.Println("Primer corrida 1-7")
  cuadrado(1,2,3,4,5,6,7)
  fmt.Println("Segunda corrida 10,20,30,40,50")
  cuadrado(10,20,30,40,50)
}
func cuadrado(n...int){
  for _, m:=range n{
    fmt.Println(m*m)
  }
}
