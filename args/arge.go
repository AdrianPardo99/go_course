package main
import ("fmt"
        "os"
        "strconv")
func main(){
  if len(os.Args)<2{
    fmt.Println("Error\nUsage: "+os.Args[0]+" <num>")
    os.Exit(1)
  }
  num,err:=strconv.Atoi(os.Args[1])
  if err==nil{
    if num%2==0{
      fmt.Println("Pair")
    }else{
      fmt.Println("Non Pair")
    }
  }else{
    fmt.Println("Error convertion")
    os.Exit(2)
  }
}
