package main
import "fmt"
func main(){
  val:=0
  val1:=1
  val2:=0
  aux:=0
  const val_c=20
  for i:=0; i<val_c; i++{
    val2=val+val1
    fmt.Println("Valor de la serie ",(i+1),": ",val)
    aux=val
    val=val2
    val1=aux

  }
}
