package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)
  // Output: HelloWorld
  
  
  fmt.Println(i,j)
  // Output: Hello World

  fmt.Printf("%s %s\n",i,j)
  // Output: Hello World
}
