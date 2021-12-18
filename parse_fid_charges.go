package main
import(
        "fmt"
        "os"
        "bufio"
        "strconv"
        "strings"
        ) 

func main(){
    var charge_file string
    if (len(os.Args) > 1 ) {
        charge_file = os.Args[1] 
    } else {
        fmt.Println("Usage: ", os.Args[0], " <charge_filename> " )
        return
    }  
    sep := "\t"
    fmt.Println("Charge file name is : ", charge_file) 
    in, err := os.Open(charge_file)
    if err != nil {
        panic(err) 
    } 
    defer in.Close()
    r := bufio.NewReader(in)
 //   var lines map[string]string
//lines = make(map[string]string )
    line, err := r.ReadString('\n')
    var total float64 = 0
    for err == nil {
        fields := strings.Split(line, sep)
        if (len(fields) > 5 ) {  
            amount := strings.TrimSpace(strings.Replace(fields[5], "$", "", -1 )) 
            dlr_amt,_ := strconv.ParseFloat(amount, 64) 
            total +=dlr_amt

        }

        line, err = r.ReadString('\n') 
    } 
    fmt.Println("Total amount :", total) 
}  
