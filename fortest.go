package main
import (
        "fmt"
        "time"
       )

func main() {
s1 :=[]int{1,2,3,4,5,6,7}
    s2 := make([]interface{}, len(s1))
    var i int
    fmt.Println("before assigning value to s2:", s2) 
    for i, s2[i] = range s1 {}
    fmt.Println(s2)
    i = 0
    i,s2[i] = 1, s1[1]
    fmt.Println(s2) 
    t0 := time.Now()
    for i :=0 ; i < 10000000; i++ {
        if i % 10000 == 0 {
            fmt.Println(i, "\n")
        } 
    }
    t1 := time.Now()
    fmt.Println("Time spent is : ",  t1.Sub(t0), " and  t0 is : ", t0 ) 
}  
