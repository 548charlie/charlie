//listdir.go
package main

import (
    "os"
    "io/ioutil"
    "fmt"
)

func ListDir(dir string) ([]os.FileInfo, error) {
    return ioutil.ReadDir(dir)
}

func main() {
    dir := "./"
    if len(os.Args) > 1 {
        dir = os.Args[1]
    }
    fi, err := ListDir(dir)
    if err != nil {
        fmt.Println("Error", err)
    }

    for _, f := range fi {
        d := "-"
        if f.IsDir() { d = "d" }
        fmt.Printf("%s %o %d %s %s\n", d, f.Mode() & 0777, f.Size(), f.ModTime().Format("Jan 2 15:04"), f.Name())
    }
}

