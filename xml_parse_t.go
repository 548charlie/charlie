package main

import (
    "encoding/xml"
    "fmt"
    "os"
)

type MyElements struct {
    XMLName  xml.Name `xml:"myElements"`
    Count    int      `xml:"count,attr"`
    Myelement []MyElement
}

type MySubStructA struct {
    Value int    `xml:"value"`
    Units string `xml:"units"`
}
type MySubStructB struct {
    XMLName xml.Name `xml:"mySubStructB"`
    Id      string   `xml:"id,attr"`
    Value   int      `xml:"value"`
    Units   string   `xml:"units"`
}

type MyElement struct {
    XMLName      xml.Name `xml:"myElement"`
    StartDate    string   `xml:"StartDate"`
    EndDate      string   `xml:"EndDate"`
    AsOfDate     string   `xml:"asOfDate"`
    numberX      float32  `xml:"numberX"`
    numberY      int      `xml:"numberY'`
    numberZ      float32  `xml:"numberZ"`
    mySubStructB []MySubStructB
    mySubStructA []MySubStructA
}

func main() {
    Fp, err := os.Open("xml_test.xml")
    defer Fp.Close()
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
    d := xml.NewDecoder(Fp)
    var myelements MyElements
    err = d.Decode(&myelements)
    if err != nil {
        fmt.Println(err)

    }
    /*
    Fp, err := os.Open("/home/francisco/Workspace/go/src/mycompany/ingestor/myElments.xml")
        if err != nil {
            fmt.Println(err)
            os.Exit(0)
        }
        FStat, err := Fp.Stat()
        bytes = make([]byte, FStat.Size())
        _, err = Fp.Read(bytes)
        if err != err {
            fmt.Println(err)
            os.Exit(0)
        }
        xml.Unmarshal(bytes, &myelements)

    */
    fmt.Println(myelements)
    fmt.Println(myelements.XMLName)
    fmt.Println("Count: %i", myelements.Count)
    fmt.Println("myelement ", myelements.Myelement) 
}

//f(t)
