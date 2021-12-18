package main
import (
        "fmt"
        "os"
        "path/filepath"
        "strings"
        "bufio"
        "strconv"
)
func compare_msg(msg1, msg2 string, count int) string{
    
    segments1 := strings.Split(msg1, "\r")
    segments2 := strings.Split(msg2, "\r")
    //id2 := 0
    if msg1 == msg2 {
        fmt.Println("Msg#", count,  "messages are same")
        return "same"
    } 
    fmt.Println("Msg#:", count) 
    if len(segments1) != len(segments2) {
        fmt.Println("Two messages have different number of segments")  
    } 
        compare_segments(segments1, segments2) 
    return "same"
}  
func compare_segments(segments1, segments2 []string){
    id2 := 0
    segs1 := make(map[string]  string)
    segs2 := make(map[string]  string ) 
    seen := make(map[string] string ) 
    for _, seg := range segments1 {
        if seg == "" {continue} 
        segid := seg[0:3]  
        if  _, ok := segs1[segid]; !ok {
            segs1[segid] = seg
        } else {
            segid = segid + strconv.Itoa(id2)
            id2++
            segs1[segid] = seg 
        }   
    } 
    id2 = 0
    for _, seg := range segments2 {
        if seg == "" {continue} 
        segid := seg[0:3]
        if  _, ok := segs2[segid]; !ok {
            segs2[segid] =  seg
        } else {
            segid = segid + strconv.Itoa(id2)
            id2++
            segs2[segid] = seg 
        }   
    }
    for seg, fields := range segs1 {
        if fields2, ok := segs2[seg]; ok {
            seen[seg]=seg 
            if fields != fields2 { 
                fmt.Println(fields)
                fmt.Println(fields2) 
                compare_fields(seg, fields, fields2)
            }
        }  
    }
    for seg, _ := range segs1 {
        if _, ok := seen[seg]; !ok {
            fmt.Println(seg, "extra segment present in file1") 
        }  
    } 
    for seg, _ := range segs2 {
        if _, ok := seen[seg]; !ok {
            fmt.Println(seg, "extra segment present in file2") 
        }  
    } 
}  

func compare_fields (seg, seg1, seg2 string) {
    
    fields1 := strings.Split(seg1, "|")
    fields2 := strings.Split(seg2, "|")
    if len(fields1) >= len(fields2) {
        for id, field := range fields1 {
            if id <= len(fields2) {
                if field != fields2[id] {
                    //fmt.Println(seg, id,field, fields2[id]) 
                    compare_subfields(seg, id, field,fields2[id] ) 
                }  
            }  
        } 
    } else {
        for id, field := range fields2 {
            if id <= len(fields1) {
                if field != fields1[id] {
                    fmt.Println(seg, id, fields1[id] , field) 
                }  
            }  
        }
    
    }    
}  

func compare_subfields(seg string, fieldid int, field1, field2 string) {
    subfields1 := strings.Split(field1, "^")
    subfields2 := strings.Split(field2, "^")
    if len(subfields1) >= len(subfields2) {
        for id, subfield := range subfields1 {
            if id <= len(subfields2) {
                if subfield != subfields2[id] {
                    fmt.Println(seg,fieldid, id+1,subfield, subfields2[id] ) 
                }  
            }  
        } 
    } else {
        for id, subfield := range subfields2 {
            if id <= len(subfields1) {
                if subfield != subfields1[id] {
                    fmt.Println(seg,fieldid, id,subfields1[id] , subfield ) 
                }  
            }  
        }
    }   
}  
func main(){
    var file1, file2, idfile string
    fmt.Println(os.Args)
    if  len(os.Args) >= 4 && len(os.Args) < 5  {
        file1 = os.Args[1]
        file2 = os.Args[2]
        idfile = os.Args[3] 
    }  else {
        fmt.Println("print Help on how to use program") 
    } 
    fmt.Println("number of args:", len(os.Args), "file1:", file1, " file2:", file2, "idfile", idfile) 
    outfile := strings.Split(filepath.Base(file1),".")[0] + "_" + strings.Split(filepath.Base(file2),".")[0] + ".csv"
    fmt.Println("outputfile:", outfile) 
    if _,err := os.Stat(file1); os.IsNotExist(err) {
        fmt.Println("No such file or directory %s", file1)
        return
    }
    if _,err :=os.Stat(file2); os.IsNotExist(err) {
        fmt.Println("No such file or directory %s", file2) 
    }
    var messages1, messages2 []string
    messages1 = make([]string, 0)
    messages2 = make([]string, 0) 
    file, err := os.Open(file1)
    if err != nil {
        panic(err) 
    } 
    defer file.Close()
    s := bufio.NewScanner(file)
    for s.Scan() {
        line := s.Text() 
        if line == "" {
            continue
        } 
        messages1 = append(messages1, line) 
        
    }  
    file, err = os.Open(file2)
    if err != nil {
        panic(err) 
    } 
    defer file.Close()
    s = bufio.NewScanner(file)
    for s.Scan() {
        line :=s.Text() 
        if line == "" {continue} 
        messages2 = append(messages2, line)  
    }
    if len(messages1) >= len(messages2) {
        for id, line := range messages1 {
            compare_msg(line, messages2[id],id+1 ) 
        } 
    }  else {
        for id, line := range messages2 {
            compare_msg(messages1[id] , line, id+1 ) 
        } 
    }  

}  
