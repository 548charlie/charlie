package main

import (
    "fmt"
    "os"
    "encoding/xml"
    "io/ioutil"
    "bufio"
)

type Book struct {
    XMLName xml.Name `xml:"book"`
    id string `xml:"book,attr"`
    Author string
    Title string
    Genre string
    Price float64
    Publish_date string
    Description string
}
func main() {
    bytes := make([]byte, 1024)
    xmlFile, err := os.Open("onebook.xml")
    defer xmlFile.Close()
    reader := bufio.NewReader(xmlFile)
    bytes, err = ioutil.ReadAll(reader)
    s := string(bytes)
    fmt.Println(s)
    if err != nil {
        fmt.Println("Error Opening file", err)
        return
    }
    var v Book  
    xml.Unmarshal(bytes, &v)
    fmt.Println(v)
}
