package main
import (
    //"encoding/xml"
    "fmt"
    "os"
) 

func main() {
    var filename string 
    if len(os.Args)  < 2 {
        fmt.Println("usage: ", os.Args[0] ," <filename>" ) 
        os.Exit(0) 
    }  else {
        filename = os.Args[1] 
    } 
    fmt.Println("filename parsing is :", filename) 
}  
