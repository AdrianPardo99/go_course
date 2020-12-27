package main
import "fmt"
type Persona struct{
  nombre string
  apellidos string
  edad uint8
  sexo string
}
func main(){
  var persona=Persona{nombre:"Juan",apellidos:"Perez",edad:18,sexo:"M"}
  persona1:=Persona{"Adrian","Gonzalez",21,"M"}
  fmt.Println(persona)
  fmt.Println(persona1)
}
