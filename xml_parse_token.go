package main

import (
    "fmt";
    "os";
    "encoding/xml"
)

func main() {
    var (
        err os.Error
        f *os.File
        p *xml.Parser
        token xml.Token
    )
    if len(os.Args) < 2 {
        fmt.Printf("Usage:", os.Args[0], "<filename>" ) 
    }  
filename := os.Args[1] 
    defer func() {
        if err != nil { fmt.Println(err) }
        f.Close()
    }()

    f, err = os.Open(filename, os.O_RDONLY, 0666);
    if err != nil { return }

    indent := 0

    p = xml.NewParser(f)
    for token, err = p.Token(); err == nil; token, err = p.Token() {
        switch t := token.(type) {
        case xml.StartElement:
            for i := 0; i < indent; i++ { fmt.Print(" ") }
            fmt.Printf("<%s>\n", t.Name.Local)
            indent++
        case xml.EndElement:
            indent--
            for i := 0; i < indent; i++ { fmt.Print(" ") }
            fmt.Printf("</%s>\n", t.Name.Local)
        }
    }
}

