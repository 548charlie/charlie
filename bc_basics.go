package main
import "fmt"

const (
        pi = 3.14
        Truth = false
        Big = 1 << 100
        small = Big >> 99
        ) 
type Artist struct {
    name, genre string
    songs int
} 

type Stringer interface {
    String() string 
}  

type fakeString struct {
    content string
} 

func (s *fakeString) String() string {
    return s.content
}   

func printString(value interface{} ) {
    switch str := value.(type) {
        case string:
            fmt.Println(str)
        case Stringer:
            fmt.Println(str.String() ) 
    }  
}  
func newRelease (a *Artist) int {
    a.songs++
    return a.songs
}  
func main() {
    const greeting = "Hello World"
    fmt.Println(greeting)
    fmt.Println(pi)
    fmt.Println(Truth) 
    fmt.Printf("%s greeting pi %.3f and Truth %t\n", greeting, pi, Truth) 

    me := &Artist{name: "Dinakar", genre : "guilter", songs: 42}
    fmt.Printf("%s released their %dth song\n", me.name, newRelease(me) )
    fmt.Printf("%s has total of %d songs\n", me.name, me.songs) 

    s := &fakeString{"Dinakar Desai, Chinchali string"}
    printString(s)
    printString("Hello, Dinakar Desai") 
}  
