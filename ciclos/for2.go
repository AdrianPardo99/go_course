package main
import ("fmt"
        "strconv"
        "os")
func main(){
  if len(os.Args)<2{
    fmt.Println("Error\nUsage "+os.Args[0]+" <number>")
    os.Exit(1)
  }
  val,err:=strconv.Atoi(os.Args[1])
  if err!=nil{
    fmt.Println("Error convertion integer")
    os.Exit(1)
  }
  arr_fib:=[]uint64{0,1}
  p:=1
  /* Ciclo repetitivo hasta 10 veces */
  for i:=0;i<val;i++{
    arr_fib=append(arr_fib,(arr_fib[p+i]+arr_fib[p-1+i]))

  }
  fmt.Println("Serie fibonacci ",arr_fib)
}
