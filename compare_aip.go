package main
import (
    "fmt"
    "os"
    "encoding/csv"
    "io"
    "strconv"
)
var help string = `===========================
Following is an example on how to use JXPORT in Epic and
what options to chose for this program to work
RELB>d main^JXPORT
INI: AIP-Interface Profile
Import spec:
Item(s): all
Put multiple response data into a single row (per record)? no
Export null data as <NULL>? No
Hide unused items? No
Include category validation? No
Export category names? Yes
Include INI validation? No
Export raw data (unpadded .1) for networked items? yes
Export networked record names in adjacent columns? yes
Export networked record private external IDs? No
Protect sheet structure? No
Profile Recs: all
Use last DAT? Yes
===================================
fetch xml file from the server
Open xml file with MS Excel program and save it as csv files.
At present output it sent to screen. You can redirect it to a file
Open the output with MS Excel. Only differences are shown.
`         


type rows struct {
    name string
    row []string
}
func parse_csv(filename string, delim rune) ( all_columns map[string][]string,  rows map[string]map[string]string,   err error){
    rows = make(map[string]map[string]string    ) 
    all_columns = make(map[string][]string  ) 
    profiles := make(map[string]string   ) 
    file, err := os.Open(filename)
    defer file.Close()

    var column_num int
    var profile string = ""
    prev, current := "one", "two"
    if err != nil {
        fmt.Println(filename, " " , err)
        return  all_columns,rows,err
    } else {
        file.Seek(0, os.SEEK_SET)
        r := csv.NewReader(file)
        r.Comma =delim
        r.TrailingComma = true
        r.TrimLeadingSpace = true
        row, err := r.Read()
        if err != nil {
            return all_columns,rows,err
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
                all_columns[prev] = row 
                profiles = nil
                profiles = make(map[string] string) 
            }  
            var column_names [] string 
            if row[0] ==  "PROFILE RECORD" {
                column_names = row
                column_name :="PROFILE VARIABLES RECORD NAME"
                all_columns["column"] = row 
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
        
    return all_columns,rows, err
}
func compare_variables(rowsa, rowsb map[string]map[string]string  ) {
    for keya, valuea := range rowsa {
        if valueb, ok := rowsb[keya]; ok {  
            for k, v := range valuea {
                if _, ok := valueb[k]; !ok {
                    if v != "" { 
                        fmt.Println(keya, ",", k, ",", v,",", "NO_VALUE") 
                    }
                }else {
                    if valuea[k] != valueb[k] {
                        fmt.Println(keya,",",k,",",v, ",",valueb[k] ) 
                    }   
                } 
            } 
        }
    }     
    for keyb, valueb := range rowsb {
        if valuea, ok := rowsa[keyb]; ok {
            for k, v := range valueb {
                if _, ok := valuea[k]; !ok {
                    if v != "" { 
                        fmt.Println(keyb, ",", k,",NO_VALUE,", v)
                    }
                }else {
                    if valuea[k] != valueb[k] {
                        fmt.Println(keyb,",",k,",",valuea[k],",",v ) 
                    }    
                }  
            } 
        }  
    } 
} 
func compare_columns(cola, colb map[string][]string  ){
column_names := cola["column"] 
    for key, valuea := range cola {
        if valueb, ok := colb[key]; ok {
            for i, vala := range valuea {
                valb := valueb[i]  
                if vala != valb {
                    if valb == "" {valb = "NO_VALUE"} 
                    if vala == "" {vala = "NO_VALUE"} 
                    temp := "ZZ_" + key
                    fmt.Println(temp,",",column_names[i],",",  vala, ",",valb,"," ) 
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
        fmt.Println(help) 
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
    all_columnsa,rows2,err := parse_csv(file1,delim )
    all_columnsb,rows3, err := parse_csv(file2, delim)
    fmt.Println("Profile Record, Profile Name,",file1, ",", file2) 
    compare_variables(rows2, rows3) 
    compare_columns(all_columnsa, all_columnsb) 
    fmt.Println(err) 
    return
}
