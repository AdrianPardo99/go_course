package main
import "fmt"
func main(){
  arr_fib:=[]int64{0,1}
  val:=10
  i:=0
  p:=1
  for i<val{
    arr_fib=append(arr_fib,(arr_fib[p+i]+arr_fib[p-1+i]))
    i++
  }
  fmt.Println("Serie fibonacci ",arr_fib)
}
