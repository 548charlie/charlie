package main
//"io"
//"strconv"
//"strings"

import(  
"fmt"
"os"
"io"
"encoding/csv"
)

func main() {
    var filename string
    if len(os.Args) < 2  {
        fmt.Println("Usage: ", os.Args[0], "filename" )
        return
    } 
    if len(os.Args) >1 {
        filename = os.Args[1] 
    } 
    fmt.Println("filename is : ", filename) 
    file, err := os.Open(filename) 
   if err != nil {
        fmt.Println("filename :", err)         
        return
    } 
    defer file.Close()
    file.Seek(0, os.SEEK_SET) 
    r := csv.NewReader(file) 
    r.Comma = '|'
    r.TrailingComma = true
    r.TrimLeadingSpace = true
    row, err := r.Read() 
    if err != nil {
        fmt.Println(err) 
        return
    } 
    for err == nil {
        if err == io.EOF {
            break
        } 
        fmt.Println("Number of columns: " , len(row), row )
        fmt.Println() 
        row, err = r.Read() 
    } 
temp := fmt.Sprintf("%-10s,%-100s,%-400s", "test1", "test2", "*") 
          fmt.Println(temp) 
}  
