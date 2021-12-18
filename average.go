package main
import "fmt"

func average(xs []float64 ) float64 {
    var x float64 = 0.0
    for _, v := range xs {
        x += v
    } 
    return x
}  

func main() {
    xs := make([]float64, 3, 5)
    xs[0]=0.2
    xs[1] =1.1
    xs[2] = 2.3 
    fmt.Printf("Average of xs %v is %f\n", xs, average(xs)) 
    fmt.Println("Hello World") 
} 
