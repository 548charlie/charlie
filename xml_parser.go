package main

import (
        "fmt";
        "os";
        "encoding/xml"
        "bufio"
        "io/ioutil"
       )
func parse_by_element(filename string) {
    f, err := os.Open(filename);
    defer f.Close()
        if err != nil {
            fmt.Println(err)
                os.Exit(1)
        }

indent := 0

            p := xml.NewDecoder(f)
            for token, err := p.Token(); err == nil; token, err = p.Token() {
                switch t := token.(type) {
                    case xml.StartElement:
elem := xml.StartElement(t)
          for i := 0; i < indent; i++ { fmt.Print(" ") }
      fmt.Printf("%s: ", elem.Name.Local)
          //attrs := elem.Attr
          //for _,v := range attrs {
          //    fmt.Println( v.Name.Local , " = ", v.Value)
          //}
          //indent++
                    case xml.EndElement:
          //indent--
          //for i := 0; i < indent; i++ { fmt.Print(" ") }
          //fmt.Printf("</%s>\n", t.Name.Local)
                    case xml.CharData:
data := string([]byte(xml.CharData(t)))
          fmt.Println(data)
                }
            }

}

func parse_unmarshal(filename string) {
    type Person struct {
        XMLName xml.Name `xml:"Person"`
            FirstName string `xml:"firstname"`
            LastName string `xml:"lastname"`
            Address string `xml:"address"`
            City string `xml:"city"`
            State string `xml:"state"`
    }
    var person  Person = Person{FirstName: ""}
    fmt.Println(person.FirstName)
        f, err := os.Open(filename)
        if err != nil {
            fmt.Println(err)
                os.Exit(1)
        }
    defer f.Close()
        reader := bufio.NewReader(f)
        allData, err := ioutil.ReadAll(reader)

        fmt.Println("raw ",string([]byte  (allData)))
        err = xml.Unmarshal(allData, &person)
        if err !=nil {
            fmt.Println(err)
        }
    fmt.Println(person)
        fmt.Println(person.FirstName)
        fmt.Println(person.Address)
        fmt.Println(person.City)
        fmt.Println(person.State)


}
func parse_excel(filename string) {
    type Name struct {
        Id int `xml:"id,attr"`
            Name string `xml:",chardata"`
    }
    type Result struct {
        Name Name `xml:"name"`
    }

    fmt.Println("==============\nparse_excel_1\n==============\n")
        Fp,err := os.Open(filename)
        if err != nil {
            fmt.Println(err)
                os.Exit(1)
        }
    defer Fp.Close()
        d := xml.NewDecoder(Fp)
        var result Result
        err = d.Decode(&result)
        fmt.Println(result)
        fmt.Println(result.Name.Id, " ", result.Name.Name)
}
func parse_excel_2(filename string) {
    type Data struct {
        Data string `xml:",chardata"`
    }
    type Result struct {
        XMLName xml.Name `xml:"row"`
            Data []Data `xml:"data"`
            Name string `xml:"name"`
    }
    fmt.Println("==============\nparse_excel_2\n==============\n")
        Fp,err := os.Open(filename)
        defer Fp.Close()
        if err != nil {
            fmt.Println(err)
                os.Exit(1)
        }
d := xml.NewDecoder(Fp)
       var result Result
       err = d.Decode(&result)
       fmt.Println(result)
       //fmt.Println(result.Name.Id, " ", result.Name.Name)

}
func parse_excel_3(filename string) {
    type Data struct {
        Data string `xml:",chardata"`
            State int `xml:"state,attr"`
    }
    type Col struct {
        //XMLName xml.Name `xml:"col"`
        Data Data `xml:"data"`
    }
    type Result struct {
        //XMLName xml.Name `xml:"row"`
        Col []Col `xml:"col"`
    }
    fmt.Println("==============\nparse_excel_3\n==============\n")
        Fp,err := os.Open(filename)
        defer Fp.Close()
        if err != nil {
            fmt.Println(err)
                os.Exit(1)
        }
d := xml.NewDecoder(Fp)
       var result Result
       err = d.Decode(&result)
       fmt.Println(result)
       for _,elem := range result.Col {
           fmt.Println("elem", elem.Data.Data)
               fmt.Println("State", elem.Data.State)
       }
   //fmt.Println(result.Name.Id, " ", result.Name.Name)

}
func parse_person(filename string) {
/*
    <Person>
        <FullName>Grace R. Emlin</FullName>
        <Company>Example Inc.</Company>
        <Email where="home">
            <Addr>gre@example.com</Addr>
        </Email>
        <Email where='work'>
            <Addr>gre@work.com</Addr>
        </Email>
        <Group>
            <Value>Friends</Value>
            <Value>Squash</Value>
        </Group>
        <City>Hanga Roa</City>
        <State>Easter Island</State>
    </Person>
*/
    type Email struct {
        Where string `xml:"where,attr"`
            Addr string
    }
    type Group struct {
        Value string `xml:",chardata"`
    }
    type Person struct {
        FullName string `xml:"FullName"`
            Company string `xml:"Company"`
            City string `xml:"City"`
            Group []Group  `xml:"Group>Value"`
            State string `xml:"State"`
            Email []Email
    }
    fmt.Println("==============\nparse_person\n==============\n")
        Fp,err := os.Open(filename)
        if err != nil {
            fmt.Println(err)
                os.Exit(1)
        }
    defer Fp.Close()
        d := xml.NewDecoder(Fp)
        var result Person
        err = d.Decode(&result)
        fmt.Println(result)
        fmt.Println("FullName: ",result.FullName)
        fmt.Println("Company: ",result.Company)
        fmt.Println("City: ",result.City)
        fmt.Println("State: ",result.State)
        fmt.Println("emails:", result.Email)
        for _,email := range result.Email {
            fmt.Println(email.Addr)
                fmt.Println(email.Where)
        }
    fmt.Println("Groups: ", result.Group)
        for _,grp := range result.Group {
            fmt.Println(grp.Value)
        }

}
func main() {
    parse_by_element("stuff.xml")
        parse_unmarshal("stuff.xml")
        parse_excel("excel.xml")
        parse_excel_2("excel_2.xml")
        parse_excel_3("excel_3.xml")
        parse_person("person.xml")
}

