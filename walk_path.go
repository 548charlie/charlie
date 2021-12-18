package main
import (
    "path/filepath"
    "os"
    "fmt"
    "flag"
    "strings"
)
func main() {

    markFn := func(path string, info os.FileInfo, err error) error {
    if strings.Contains(path, "src")  {
        fmt.Println("We are skipping this directory ", path)
        return filepath.SkipDir
    }
    if err != nil {
        return err
    }
    fmt.Println(path)
    return nil
}
    fmt.Println(flag.NArg() )
    path := flag.Arg(0)
    fmt.Println("path ", path)
    pwd, err := os.Getwd()
    if err == nil {
        fmt.Println("pwd ", pwd)
    }
    root :="c:/dinakar/go_test"

    err = filepath.Walk(root, markFn)
    if err != nil {
        panic(err)
    }

}
