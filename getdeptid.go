package main
import (
    "fmt"        
    "os"
    "bufio"
    "regexp"
)

func main() {
    fmt.Println("Hello World")
    var pattern string
    if len(os.Args) > 1  {
       pattern = os.Args[1] 
    } else {
        fmt.Println("Please provide the filename to the" + os.Args[0])
        return
    }
    inFile, _ := os.Open("c:/junk/dept_codes.txt")
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
    reg,_ := regexp.Compile(pattern)
  for scanner.Scan() {
    //fmt.Println(scanner.Text())
    line := scanner.Text()
//              fmt.Println(line)
   
    matched := reg.Match([]byte(line))
     if matched  {
        fmt.Println("Found " + line)
     }
  
  }
}
