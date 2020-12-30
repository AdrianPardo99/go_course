package main
import ("fmt"
        "io/ioutil"
        "os")
func main(){
  if len(os.Args)<2{
    fmt.Println("Error\nUsage: "+os.Args[0]+" <file>")
    os.Exit(1)
  }
  name:=os.Args[1]
  archivo,err:=ioutil.ReadFile(name)
  catch_error(err)
  fmt.Println("Lectura desde archivo\n\tContenido:")
  fmt.Println(string(archivo))

}
func catch_error(e error){
  if e!=nil{
    panic(e)
    os.Exit(2)
  }
}
