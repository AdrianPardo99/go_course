package main
import ("fmt"
        "time")
func main(){
  enum_dia:=time.Now()
  hoy:=enum_dia.Weekday()
  switch hoy {
  case 0:
    fmt.Println("Domingo")
  break
  case 1:
    fmt.Println("Lunes")
  break
  case 2:
    fmt.Println("Martes")
  break
  default:
    fmt.Println("Otro día")
  }
}
