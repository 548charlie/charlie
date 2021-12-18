package main
import ("fmt"
        "time"
        )


func Factorial(n int ) (result int) {
    if n == 0 {
        result = 1
    } else {
        result= n * Factorial(n -1 )
    }
    return

}
func nFactorial (x int ) int {
    if x == 0 {
        return 1
    } else {
        return x * nFactorial(x - 1)
    }
    return x
}
func main(){
    fmt.Printf("Hello World\n")
    fmt.Println("Time now is :", time.Now() ) 
    list := []string{"a", "b", "c","d","e"}
    for k, v := range list {
        if k > 1 && k < 4 {
            continue
        }
        fmt.Printf("%d  %s\n", k, v)
    }
    s := "This is test line"
    for pos, char := range s {
        fmt.Printf("%d == %c\n", pos, char)
    }
    /* declaration -- monthdays :=make(map[string]int ) */
    monthdays := map[string]int {
               "Jan":31, "Feb": 28, "Mar": 31,
                "Apr":30, "May":31, "Jun":30,
                "Jul":31, "Aug":31, "Sept":30,
                "Cot":31, "Nov":30, "Dec":31,
               }
    year := 0
    for _, days := range monthdays {
        year += days
    }
    fmt.Printf("Number of days in a year: %d\n", year)
    fmt.Println("Factorial of 10 is : ", Factorial(10)  )
    fmt.Println("Factorial of 10 is : ", nFactorial(10) )
}
