package main

import "fmt"

func main() {
   fmt.Println("Hello, World")
   var name string
   name = ""
   for i := 0; i < 10; i++ {
      name = name + "A"
      fmt.Printf("%v\n", name)
   }
}
