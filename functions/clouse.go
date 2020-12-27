package main
import "fmt"
func main(){
  fmt.Print("Precio de $10\n")
  fmt.Println(calc(10))
}
func calc(n float32)(string,float32){
  iva:=func()float32{
    return n*1.16
  }
  return "Precio con iva: $",iva()
}
