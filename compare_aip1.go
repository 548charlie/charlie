package main
import (
    "fmt"
    "os"
    "encoding/csv"
    "io"
    "strconv"
)

type rows struct {
    name string
    row []string
}
func parse_csv(filename string, delim rune) ( rows map[string]map[string]string,   err error){
    rows = make(map[string]map[string]string    ) 
    profiles := make(map[string]string   ) 
    file, err := os.Open(filename)
    defer file.Close()

    var column_num int
    var profile string = ""
    prev, current := "one", "two"
    if err != nil {
        fmt.Println(filename, " " , err)
        return rows,err
    } else {
        file.Seek(0, os.SEEK_SET)
        r := csv.NewReader(file)
        r.Comma =delim
        r.TrailingComma = true
        r.TrimLeadingSpace = true
        row, err := r.Read()
        if err != nil {
            return rows,err
        }
        for err == nil {
            if err == io.EOF {
                break
            }
            if row[0]  != "" {
                profile = row[0] 
            }  
            if prev != profile {
                prev = current
                current = profile
//                fmt.Println("prev ", prev, " current ", current, " profile ", profile) 
                rows[prev]=profiles 
                profiles = nil
                profiles = make(map[string] string) 
            }  
            var column_names [] string 
            if row[0] ==  "PROFILE RECORD" {
                column_names = row
                column_name :="PROFILE VARIABLES RECORD NAME"
                for i, value := range column_names {
                    if value == column_name {
                        column_num = i
                    } 
                } 
            } 
            
            //fmt.Println("column 1", profile, " profile variable ", row[column_num], "-", row[column_num +1]   )  
            //value := row[column_num] + "|" + row[column_num +1]  
            profiles[row[column_num] ] = row[column_num +1] 
            //fmt.Println("Number of columns ", len(row), row)
            row, err = r.Read()

        }
    }
    return rows, err
}
func compare_variables(rowsa, rowsb map[string]map[string]string  ) {
    if len(rowsa) >= len(rowsb) {
        for keya, valuea := range rowsa {
            fmt.Println("key : ", keya, " values" , valuea) 
            if valueb, ok := rowsb[keya]; ok {
                if len(valuea) >= len(valueb) {
                    for k, v := range valuea {
                        if _, ok := valueb[k]; !ok {
                            fmt.Println("\"",keya,"\",\"",k,"\",\"",v,"\",No Value") 
                        }  
                    }  
                } else {
                    for k, v := range valueb {
                        if _, ok := valuea[k]; !ok {
                            fmt.Println("\"",keya,"\",\"",k,"\",NoValue\"",v,"\"")  
                        }   
                    } 
                }    
            }  
        } 
    } else {
        for keyb, valueb := range rowsb {
            //fmt.Println("key : ", keyb, " values" , valueb) 
            if valuea, ok := rowsa[keyb]; ok {
                if len(valuea) >= len(valueb) {
                    for k, v := range valuea {
                        if _, ok := valueb[k]; !ok {
                            fmt.Println("\"",keyb,"\",\"",k, "\",\"",v, "\", NoValue") 
                        }  
                    } 
                } else {
                    for k, v := range valueb {
                        if _, ok := valuea[k]; !ok {
                            fmt.Println("\"",keyb,"\",\"",k, "\",NO Value,\"",v,"\"") 
                        }  
                    } 
                } 
            }  
        } 
    }    
} 
func main(){
    var file1, file2 string
    if len (os.Args)< 3 {
        fmt.Println("Usage: ", os.Args[0], "file1.csv file2.csv [delimiter] " )
        fmt.Println("Both files are csv files and delimiter is comma(,) delimiter is optional parameter\n if given it will be used to parse the csv file" )
        return
    }   
    if len(os.Args)  >= 3 {
        file1 =os.Args[1]
        file2 = os.Args[2]  
    }

    delim := ','
    if len(os.Args) >= 4 {
        file1 =os.Args[1]
        file2 = os.Args[2] 
        var sep string
        sep,_ = strconv.Unquote(`'` + os.Args[3]+`'` )
        delim = ([]rune(sep))[0]
    
    }  
    rows2,err := parse_csv(file1,delim )
    rows3, err := parse_csv(file2, delim)
    fmt.Println("Profile Record, Profile Name,",file1, ",", file2) 
    compare_variables(rows2, rows3) 
    fmt.Println(err) 
    return
}
