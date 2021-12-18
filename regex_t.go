package main
import (
    "fmt"
    "regexp"
) 

func main() {
    t := "This is test line to be 123435 713-568-7854 used with regexp"
    fmt.Println(t) 
    match, err := regexp.MatchString("\\d+", t)
    if err != nil {
        fmt.Println(err) 
    } 
    if match {
        fmt.Println("matched") 
    }
    var match_to string =`(\d{2,}-\d{1,}-\d{3,})|(used)`
    reg, err := regexp.Compile(match_to)
    if err != nil {
        fmt.Println(err) 
    } 
    matched := reg.Find([]byte(t)  ) 
    if len(matched) > 0  { 
        fmt.Println("Find:" ,string([]byte(matched))) 
    } else {
        fmt.Println("Did not match") 
    }
    stMatches:= reg.FindAllString(t, 5)
    fmt.Println("FindAllString:" ,stMatches) 
    reg1,_ := regexp.Compile(`([1-9]1\d+)|(used)`) 
    subMatch := reg1.FindStringSubmatch(t)
    fmt.Println("FindStringSubmatch: ",subMatch) 
}  
