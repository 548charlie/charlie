package main
import (
    "os"
    "flag"
    "fmt"
    "bufio"
    "io"
)
var omitNewline = flag.Bool("n", false, "no final newline")
const (
    Space=""
    Newline="\n"
)
func readString(filename string) {
    f,err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    r := bufio.NewReader(f)
    line, err := r.ReadString('\n')
    for err == nil {
        fmt.Print(line)
        line, err = r.ReadString('\n')
    }
    if err != io.EOF  {
        fmt.Println(err)
        return
    }
}
func main(){
    flag.Parse()
    var s string = ""
    fmt.Println("numb of args :",flag.NArg() )
    for i:= 0; i < flag.NArg();i++ {
        fmt.Println(flag.Arg(i) )
        if i > 0 {
            s += Space
        }
        s += flag.Arg(i)
    }
    if ! *omitNewline {
        s+=Newline
    }
    os.Stdout.WriteString(s)
    readString("cmdline.go")
}
