package main
import (
        "fmt"
        "os"
        "strconv"
)

func fibonacci() func() int {
    n, ok := 0, false
    a, b := 0, 0
    return func() int {
        var ret int = a + b
        if false != ok {
            if  a > b {
                b += a
                ret = b
            } else {
                a += b
                ret = a
            }
        } else {
            ret = n
            n++
            if 1 < n {
                ok = true
                a = 1
            }
        }
        return ret
    }

}

func main() {
    n  := 10
    fmt.Println(os.Args[1] ) 
    if len(os.Args) > 1 {
        n,_ = strconv.Atoi(os.Args[1] )
    }  
    fmt.Println("value of args ", n) 
    f := fibonacci()
    for i := 0; i < n; i++ {
        fmt.Println(f())
    }
}
