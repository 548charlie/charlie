package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "encoding/csv"
    "io"
)

type values struct {
    name string
    value int
}
func parse_csv(filename string) {
    file, err := os.Open(filename)
    file.Seek(0, os.SEEK_SET)
    defer file.Close()
    if err != nil {
        panic(err)
    } else {
        r := csv.NewReader(file)
        r.Comma ='|'
        r.TrailingComma = true
        row, err := r.Read()
        fmt.Println(row)
        if err != nil {
            panic(err)
        }
        for err == nil {
            if err == io.EOF {
                break
            }

            fmt.Println("Number of columns ", len(row), row)
            row, err = r.Read()

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
    parse_csv(filename)
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
