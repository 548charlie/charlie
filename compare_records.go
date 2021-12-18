package main
import (
    "fmt"
    "os"
    "encoding/csv"
    "io"
    "strconv"
    "strings"
    "path/filepath"
    "regexp"
)
var help string = `===========================
Following is an example on how to use JXPORT in Epic and
what options to choose for this program to work
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
At present output it sent to the file written to current directory. 
Make sure you have write permissions in the current directory.
Open the output with MS Excel. Only differences are shown.
====================================
Author: Dinakar Desai, Ph.D.
Date: May 30, 2013
Updated on : Dec 23, 2014
If you experience any problems with the program, please contact above author.
`         


type rows struct {
    name string
    row []string
}
func parse_csv(filename string, delim rune,header_id,column_name string, offsets []int) ( all_columns map[string][]string,  rows map[string]map[string]string,   err error){
    rows = make(map[string]map[string]string    ) 
    all_columns = make(map[string][]string  ) 
    profiles := make(map[string]string   ) 
    file, err := os.Open(filename)
    defer file.Close()
    toffsets := offsets
    var column_num int
    var title string = ""
    prev, current := "one", "two"
    var  save int = 0
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
        prev = row[0]  
        for err == nil  {
            if row[0] ==  header_id {
                all_columns["column"] = row 
                for i, value := range row {
                    if value == column_name {
                        column_num = i
                    } 
                } 
            }
            if save == 0  && row[0] == "" {
                temp := row[column_num + offsets[0]] 
                if len(offsets)> 1 {  
                    offsets = offsets[1:] 
                    for _, offset := range offsets {
                        temp = temp + "_" + row[column_num + offset] 
                    } 
                }
                offsets = toffsets
                
                profiles[row[column_num]] = temp  
            } 

            if save == 1 {
                if _,ok := rows[prev];!ok {
                    rows[prev]= profiles 
                } 
                profiles = nil
                profiles = make(map[string]string )
                temp :=row[column_num + offsets[0]]
                if len(offsets) > 1 {  
                    offsets = offsets[1:] 
                    for _, offset := range offsets {
                        temp = temp + "_" + row[column_num + offset] 
                    } 
                }
                profiles[row[column_num]] = temp 
                profiles["NAME"] =title 
                save = 0
                offsets =toffsets
            } 

            row, err = r.Read()
            if err == io.EOF {
                
                save = 1
                if _,ok := rows[current];!ok {
                    rows[current]= profiles 
                } 
                break

            }   
            if row[0]  != "" {
                all_columns[row[0]]=row 
                title=row[1] 
                prev = current
                current = row[0]
                save = 1
            } 
        }
    }
    return all_columns,rows, err
}
func compare_variables(rowsa, rowsb map[string]map[string]string ,out io.Writer, filter string ) {
    re := regexp.MustCompile(filter) 
    seen := make(map[string]string) 
    writer := csv.NewWriter(out) 
    var name string
    for keya, valuea := range rowsa {
        name = valuea["NAME"] 
        if re.MatchString(keya) {   
            if valueb, ok := rowsb[keya]; ok {  
                for k, v := range valuea {
                    tkey := keya +"_" + k
                    seen[tkey] = tkey 
                    if _, ok := valueb[k]; !ok {
                        if v != "" { 
                            tmp :=[]string {keya, name,"profile-"+k, v, "NO_VALUE"}
                            err := writer.Write(tmp)
                            if err != nil {
                                panic(err) 
                            } 
                            //fmt.Println(keya, ",", name,",",k, ",", v,",", "NO_VALUE") 
                        }
                    }else {
                        if valuea[k] != valueb[k] {
                            tmp := []string {keya, name,"profile-"+k, v, valueb[k] }  
                            err := writer.Write(tmp)
                            if err != nil {
                                panic(err) 
                            } 
                            //fmt.Println(keya,",",name, ",",k,",",v, ",",valueb[k] ) 
                        }   
                    } 
                } 
            } else {
                tmp := []string{keya, name,"-","-","does not exist"}  
                err := writer.Write(tmp)
                if err != nil {
                    panic(err) 
                } 
            } 
        }
    }     
    for keyb, valueb := range rowsb {
        name = valueb["NAME"] 
        if re.MatchString(keyb) {  
            if valuea, ok := rowsa[keyb]; ok {
                for k, v := range valueb {
                    tkey := keyb + "_" +k
                    if _, ok = seen[tkey]; !ok  { 
                        if _, ok := valuea[k]; !ok {
                            if v != "" { 
                                tmp :=[]string {keyb, name,"profile-"+k, "NO_VALUE", v}
                                err :=writer.Write(tmp)
                                if err != nil {
                                    panic(err) 
                                } 
                                //fmt.Println(keyb, ",", k,",NO_VALUE,", v)
                            }
                        }else {
                            if valuea[k] != valueb[k] {
                                tmp :=[]string {keyb, name,"profile-"+k, valuea[k] , v}
                                err :=writer.Write(tmp)
                                if err != nil {
                                    panic(err) 
                                } 
                                //fmt.Println(keyb,",",name,",",k,",",valuea[k],",",v ) 
                            }    
                        }
                    }
                } 
            } else {
                tmp := []string {keyb, name, "-","does not exist","-"}
                err := writer.Write(tmp)
                if err != nil {
                    panic(err) 
                } 
            } 
        }
    } 
    writer.Flush() 
} 
func compare_columns(cola, colb map[string][]string, out io.Writer, filter string  ){
    writer := csv.NewWriter(out) 
    re :=regexp.MustCompile(filter) 
    //column_names := cola["column"] 
    column_namea :=cola["column"]
    column_nameb := colb["column"] 
    col_vala := make(map[string]  string )
    col_valb := make(map[string]  string )
    for key, valuea := range cola {
        for i, name := range column_namea {
            key1 := key + ":" + valuea[1] + ":" + name
            col_vala[key1] = valuea[i] 
        } 
    } 
    for key, valueb := range colb {
        for i, name := range column_nameb {
            key1 := key + ":" + valueb[1]  + ":" + name
            col_valb[key1] = valueb[i]  
        } 
    } 
    seen := make(map[string]string ) 
    
    for key, valuea := range col_vala {
        if re.MatchString(key) {  
            if valueb, ok := col_valb[key]; ok {
                    if valuea != valueb {
                        if valueb == "" {valueb = "NO_VALUE"} 
                        if valuea == "" {valuea = "NO_VALUE"} 
                        tkey :=key
                        if _, ok := seen[tkey]; !ok {  
                            recs := strings.Split(key,":") 
                            tmp := []string {recs[0], recs[1] ,recs[2],valuea, valueb }  
                            err := writer.Write(tmp) 
                            if err != nil {
                                panic(err) 
                            } 
                        } else {
                            seen[tkey]= tkey 
                        } 
                        //fmt.Println(temp,",",column_names[i],",",  vala, ",",valb,"," ) 
                    }  
            }  else {
                recs := strings.Split(key,":") 
                valueb = "NO_VALUE"
                if valuea != "" { 
                    tmp := []string {recs[0], recs[1] ,recs[2] ,valuea, valueb }  
                    err := writer.Write(tmp) 
                    if err != nil {
                        panic(err) 
                    }
                }
            }   
        }
    }
    for key, valueb := range col_valb {
        if re.MatchString(key) {  
            if valuea, ok := col_vala[key]; ok {
                    if valuea != valueb {
                        if valueb == "" {valueb = "NO_VALUE"} 
                        if valuea == "" {valuea = "NO_VALUE"} 
                        tkey :=key
                        if _, ok := seen[tkey]; !ok {  
                            recs := strings.Split(key,":") 
                            tmp := []string {recs[0] , recs[1] ,recs[2],valuea, valueb }  
                            err := writer.Write(tmp) 
                            if err != nil {
                                panic(err) 
                            } 
                        } else {
                            seen[tkey]= tkey 
                        } 
                        //fmt.Println(temp,",",column_names[i],",",  vala, ",",valb,"," ) 
                    }  
            } else {
                recs := strings.Split(key,":") 
                valuea = "NO_VALUE"
                if valueb != "" { 
                    tmp := []string {recs[0]+" column" , recs[1] ,recs[2],valuea, valueb }  
                    err := writer.Write(tmp) 
                    if err != nil {
                        panic(err) 
                    }
                }
            }   
        }
    }
    writer.Flush() 
}  
func main(){
    var file1, file2 ,ini_type , filter string
    record_types := "(AIP|AIF|WQF|CER|IIT|LLB)" 
    filter=".*"
    if len (os.Args)< 4 {
        fmt.Println("Usage: ", os.Args[0], "file1.csv file2.csv INI_type", record_types,"what to keep(regex) ", " [delimiter] " )
        fmt.Println("Example:", os.Args[0], "relb_aip.csv rel_aip.csv AIP \"940.*\"" ) 
        fmt.Println("Both files are csv files, in the above example the program will display only records that start with 940.\n One could use any regular expression to filter the records to keep. Default delimiter is comma(,) delimiter is optional parameter\n if given it will be used to parse the csv file" )
        fmt.Println(help) 
        return
    }   
    if len(os.Args)  >= 4 {
        file1 =os.Args[1]
        file2 = os.Args[2]  
        ini_type = os.Args[3] 
    }

    delim := ','
    if len(os.Args) >= 5 {
        file1 =os.Args[1]
        file2 = os.Args[2] 
        ini_type = os.Args[3] 
        filter = os.Args[4] 
    }  
    if len(os.Args) >= 6 {
        file1 = os.Args[1]
        file2 = os.Args[2]
        ini_type = os.Args[3]
        filter = os.Args[4]
        var sep string
        sep,_ = strconv.Unquote(`'` + os.Args[4]+`'` )
        delim = ([]rune(sep))[0]
    }  
    fmt.Println("filter is :", filter )  
    var header_id, column_name string
    offsets := make([]int,1, 10)  
    if ini_type == "AIP" {
       header_id = "PROFILE RECORD" 
       column_name = "PROFILE VARIABLES RECORD NAME"
       offsets[0]  = 1
    } else if ini_type == "AIF" {
        header_id = "TABLE ID"
        column_name = "VALUE"
        offsets[0]  = 2
    } else if ini_type =="WQF" {
        header_id = "WORKQUEUE ID"
        column_name = "RULE - ID RECORD NAME"
        offsets[0]  = 1
        offsets = append(offsets,2) 
    } else if ini_type == "CER" {
        header_id = "RULE ID"
        column_name = "RULE - LHS PROPERTIES RECORD NAME"
        offsets[0]  = 7
        offsets = append(offsets,9)
    } else if ini_type == "LLB" {
        header_id = "LAB ID"
        column_name = "INTERNAL MAPPING ITEM RECORD NAME"
        offsets[0] = 1 
    } else if ini_type == "IIT" {
        header_id = "TYPE ID"
        column_name = "APPLIES TO MASTERFILE"
        offsets[0] = 1 
    } else {
        fmt.Println("Usage: ", os.Args[0], "file1.csv file2.csv INI_type",record_types, " [delimiter] " )
        fmt.Println("Example:", os.Args[0], "relb_aip.csv rel_aip.csv AIP" ) 
        fmt.Println("Both files are csv files and default delimiter is comma(,) delimiter is optional parameter\n if given it will be used to parse the csv file" )
        fmt.Println(help) 
        return
    }  
    outfile := ini_type + "_"+strings.Split(filepath.Base(file1), ".")[0] + "_" +strings.Split(filepath.Base(file2), ".")[0] + ".csv"    
    out,err := os.OpenFile(outfile , os.O_RDWR | os.O_CREATE, 0666) 
    if err != nil {
        fmt.Println(err)
        return
    } 
    defer out.Close() 
    writer := csv.NewWriter(out)
    tmp :=[]string {"RECORD ID", "RECORD NAME","PROFILE NAME", file1, file2}
    err = writer.Write(tmp)  
    writer.Flush() 
    if err != nil {
        panic(err)
        return
    } 
    fmt.Println("name of output file ", outfile) 
    all_columnsa,rows2,_ := parse_csv(file1,delim,header_id,column_name, offsets )
    all_columnsb,rows3, _ := parse_csv(file2, delim,header_id, column_name, offsets)
    compare_variables(rows2, rows3, out, filter) 
    compare_columns(all_columnsa, all_columnsb,out, filter) 
    //fmt.Println(err) 
    return
}
