package main
import "fmt"
func main(){
  num1:=10.0
  num2:=4.5
  fmt.Println("Calculadora de num1 y num2")
  fmt.Println(calc(num1,num2))
}
func calc(n1 float64,n2 float64)(r1 float64,r2 float64,r3 float64,r4 float64){
  return (n1+n2),(n1-n2),(n1*n2),(n1/n2)
}
