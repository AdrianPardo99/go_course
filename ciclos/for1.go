package main
import ("fmt")
func main(){
  arr_fib:=[]uint64{0,1}
  val:=10
  p:=1
  /* Ciclo repetitivo hasta 10 veces */
  for i:=0;i<val;i++{
    arr_fib=append(arr_fib,(arr_fib[p+i]+arr_fib[p-1+i]))

  }
  fmt.Println("Serie fibonacci ",arr_fib)
}
