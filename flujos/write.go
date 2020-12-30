package main
import ("fmt"
        "os")
func main(){
  if len(os.Args)<3{
    fmt.Println("Error\nUsage: "+os.Args[0]+" <file> <arguments>...<write>")
  }
  file:=os.Args[1]
  /* Crea archivo y concatena */
  fmt.Println("Archivo: "+file)
  archivo,err:=os.OpenFile(file,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
  catch_error(err)
  for i:=2;i<len(os.Args);i++{
    write,err:=archivo.WriteString(os.Args[i]+"\n")
    catch_error(err)
    fmt.Println("Se escribio ",write)
  }
  archivo.Close()
}
func catch_error(e error){
  if e!=nil{
    panic(e)
    os.Exit(2)
  }
}
