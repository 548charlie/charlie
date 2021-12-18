package main

import ("fmt"
        "time"
        "encoding/xml"
)

type Person struct {
    XMLName xml.Name `xml:"person"`
    FirstName string `xml:"firstName"`
    MiddleName []string `xml:"middleName"`
    LastName string `xml:"LastName"`
    Age int64 `xml:"age"`
    Skills []Skill `xml:"skills,omitempty"` 
} 

type Skill struct {
    XMLName xml.Name `xml:"skill"`
    Name    string `xml:"skillName"`
    YearPracticed int64 `xml:"practice"`
}  

func main() {
    fmt.Println("hello World", time.Now() )
    var p Person
    p.FirstName ="Dinakar"
    p.MiddleName = make([]string,0 ) 
    p.MiddleName = append(p.MiddleName, "s") 
    p. MiddleName = append(p.MiddleName, "d") 
    p.LastName = "Desai"
    p.Age = 10
    p.Skills = make([]Skill, 0 )
    s := Skill{Name:"junk", YearPracticed:3} 
    p.Skills = append(p.Skills, s) 
    if xmlstring, err := xml.MarshalIndent(p,"","   "); err == nil {
        xmlstring = []byte(xml.Header + string(xmlstring) )
        fmt.Printf("%s\n", xmlstring) 
    }  
}  

