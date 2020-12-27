package main
import "fmt"

func main(){
  integer:=100
  var entero_sin_sig uint16=6
  fmt.Print("Valores trabajados\ninterger= ",integer,
    "\nentero_sin_sig= ",entero_sin_sig,"\n")
  fmt.Printf("Tipo integer= %T\nTipo entero_sin_sig= %T\n",
    integer,entero_sin_sig)
}
