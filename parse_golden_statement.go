package main
import (
        "fmt"
        "os"
        "bufio"
        "strings"
        "strconv" 
        "encoding/csv"
        "io"
       )

type values struct {
    name string
        value int
}
func parse_csv(filename string, delim rune) {
    file, err := os.Open(filename)
        file.Seek(0, os.SEEK_SET)
        defer file.Close()
        if err != nil {
            panic(err)
        } else {
            reader := csv.NewReader(file)
            reader.TrailingComma = true
            reader.Comma = delim
            total :=0.0
            for {
                record, err := reader.Read()
                if err == io.EOF {
                    break
                } else if err != nil {
                    fmt.Println("Error:", err)
                        return
                }
                fmt.Println(record) 
                fmt.Println(len(record) ) 
                if len(record) >= 3 {  
                    amount, _:= strconv.ParseFloat(record[4],32 )
                    total += amount
                    fmt.Printf("%-10s %-25s %-10s %8.2f\n",record[0],record[3], record[4], total) 
                }
            }

        }
}
func main(){
    var filename string

        if len(os.Args)  >= 2 {
            filename =os.Args[1]
        } else {
            filename = "junk1.csv"
        }
        fmt.Println(filename)
        delim := ','
        if  len(os.Args) >= 3 {
            var sep string
            sep,_ = strconv.Unquote(`'` + os.Args[2]+`'` )
            delim = ([]rune(sep))[0]  
    
        }  
        parse_csv(filename,delim)
        return
        file, err := os.Open("junk1.csv")
        file.Seek(0, os.SEEK_SET)
        defer file.Close()
        count := 1
        if err != nil {
            panic(err )
        } else {
r := bufio.NewReader(file)
       line, err := r.ReadString('\n')
       for err == nil {
           var columns []string
               columns = strings.Split(line, "|")
               fmt.Println(count, " ", columns, "number of columns", len(columns) )
               line, err = r.ReadString('\n')
               count++;
       }
        }
    fmt.Println("Total number of lines", count)
}
