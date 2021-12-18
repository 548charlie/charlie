package main
import "fmt"

func PrintNumbers (start, count int) {
    for i := 0; i < count; i++ {
        fmt.Printf("%d\n", start+i)
    }
}
func numberGen(start, count int, out chan <- int) {
    for i := 0 ; i < count; i++ {
        out <- start +i
    } 
    close(out) 
}  

func PrintNumber(in <-chan int, done chan<- bool) {
    for num := range in {
        fmt.Printf("%d\n", num) 
    } 
    done <-true
}  
func main() {
    numberChan := make (chan int)
    done := make(chan bool)
    go numberGen(1, 1000, numberChan)
    go PrintNumber(numberChan, done)
    <-done
}
