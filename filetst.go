package main
import (
    "path"
    "path/filepath"
    "fmt"
)

func main() {
    name := "c:/dinakar/go_test/filetst.go"
    fmt.Println("name of file is :", name)
    dir, file := path.Split(name)
    ext := path.Ext(file)
    fmt.Println("dir :", dir , " file name ", file)
    fmt.Println("extension of the file is :" , ext)
    globlst := dir +"*"

    fmt.Println("<",globlst, ">")
    filelist, err := filepath.Glob(globlst)
    if err != nil {
        panic(err)
    }
    fmt.Println("file list", filelist)
    for _, file := range filelist {
        fmt.Println("file :", file) 
    }
}
