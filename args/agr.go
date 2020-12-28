package main
import ("fmt"
        "os"
        "strconv")
func main(){
  if len(os.Args)<3{
    fmt.Println("Error\nUsage "+os.Args[0]+" <name> <age>")
    os.Exit(1)
  }
  fmt.Println("Hello "+os.Args[1])
  age,err:=strconv.Atoi(os.Args[2])
  if err==nil{
    if age>=18{
      fmt.Println("Welcome")
    }else{
      fmt.Println("Get out")
    }
  }else{
    fmt.Println("Error conversion")
    os.Exit(2)
  }
}
