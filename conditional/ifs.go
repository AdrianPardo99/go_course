package main
import "fmt"
func main(){
  val:=true
  if val{
    fmt.Println("El valor es verdadero")
  }else{
    fmt.Println("El valor es falso")
  }
  /* Forma encadenada */
  name:="el amor de tu vida"
  if name=="el amor de tu vida"{
    fmt.Println("Ulalalala")
  }else if len(name)>0{
    fmt.Println("Hola ",name)
  }else{
    fmt.Println("Hola extraño")
  }
}
