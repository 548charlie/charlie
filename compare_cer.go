package main
import (
    "fmt"
    "encoding/xml"
    "os"
    "strings"
    "bufio"
    "io"
    "strconv"
)
type Row struct {
    Cell Cell
} 
type Cell struct {
    Data string
}
func help() {
    help :=` 
    This program compares two JXPORTs of CER records. The JXPORT file is xml file.
    These files are created in the same way between environments or during different time span.
    In otherwords, they are comparable. This program first converts xml file into csv (pipe delimited)
    file. Then it compares record based on the Rule ID and comparison is by field.
    OutPut: csv file with the option chosen when you run the program or default is pipe (|)
    Parameters:
    1. first xml file (file extension has to be .xml and in small letters)
    2. second xml file (file extension has to be .xml and in small letters)
    3. Optional parameter, field separator in the final output. By default, it uses pipe(|) but if the
        parameter is provided like comma(,) or some other delimiter, it will use it. It is better to use
        pipe as it is not very common character found in xml files.
    `
    fmt.Println(help) 
}  
func xml2csv (filename string) {
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }

    defer f.Close()
    csvfilename := strings.Replace(filename, ".xml", ".csv", 1)
    file, err := os.OpenFile(csvfilename, os.O_CREATE|os.O_RDWR, 0666) 
    defer file.Close() 
    if err != nil {
        panic(err) 
    } 
    /*
    var row Row
    r := bufio.NewReader(f) 
    allData, err := ioutil.ReadAll(r ) 
    if err != nil {
        fmt.Println(err) 
    } 
    fmt.Println(string([]byte(allData))) 
    err = xml.Unmarshal(allData, &row) 
    if err != nil {
        fmt.Println(err) 
    } 
    fmt.Println(row)
    */
    decoder := xml.NewDecoder(f)
    var cols []string = make([]string,0)  
    var row int = 0
    var cell int = 0
    var data int = 0
    var colnum int = 0
    var ruleid string;
    var rulename string;
    for token, err := decoder.Token(); err != io.EOF; token, err = decoder.Token()  {
        if err != nil { 
            fmt.Println("File name ",filename, " has an following error:\n", err ,"\n\nPlease fix it and run again. The error is due to illegal character " )
            fmt.Println("Please run the following perl one liner from command line")
            fmt.Println("perl -pi.bak -e \"$_ =~ s/&#.*?;/ /g\"", filename) 
            os.Exit(0) 
        }
        switch t := token.(type) {
            case xml.StartElement :
                elem := xml.StartElement(t)
                if elem.Name.Local == "Row" {
                    //fmt.Println("<",elem.Name.Local ,">" )
                    row = 1
                    colnum = 0
                }
                if elem.Name.Local == "Cell" {
                    cell = 1
                    colnum++
                }
                if elem.Name.Local == "Data" {
                    data = 1
                } 
            case xml.CharData:
                if row  == 1 && cell == 1 && data == 1 { 
                    bytes := xml.CharData(t)
                    str := string([]byte(bytes)  ) 
                    str = strings.Replace(str, "\n", " ", -1) 
                    str = strings.Replace(str, "\r", " ", -1) 
                    str = strings.Trim(str,"\n")
                    str = strings.Trim(str, " ") 
                    //fmt.Printf("-%s| ",data)
                    if colnum == 1 && str != "" {
                        ruleid = str
                    } 
                    if colnum == 2 && str != "" {
                        rulename = str
                    }
                    cols = append(cols, str)
                } 
            case xml.Comment:
                fmt.Print("Comment", string(token.(xml.Comment) ) )
            case xml.EndElement :
                elem := xml.EndElement(t)
                if elem.Name.Local == "Row" {
                    row = 0
                    colnum = 0
                    //fmt.Println("") 
                    //fmt.Println("</", elem.Name.Local, ">" )
                    //fmt.Println("number of coln ", len(cols), "   " , coln ) 

                    for i := 0; i < len(cols) ; i++ {
                        colvalue := strings.Trim(cols[i], " ")  

                        if  i == 0 && colvalue == "-" {
                            colvalue = ruleid
                        } 
                        if i == 1 && colvalue == "-" {
                            colvalue = rulename
                        } 
                        colvalue = colvalue+"|" 
                         //fmt.Printf("%s", col )
                        _, err := file.Write([]byte(colvalue)  ) 
                        if err != nil {
                            panic(err) 
                        } 

                    }
                    //fmt.Println("")
                    file.Write([]byte("\n"))      
                    //cer := cols[2] +"_" +cols[4]
                    cols = make([]string,0)   
                }
                if elem.Name.Local == "Cell" {
                    cell = 0
                    if data == 0 {
                        cols = append(cols, "-") 
                    } 
                    data = 0
                }
                if elem.Name.Local == "Data" {
                    cell = 0
                } 
            default:
        }
    }
} 
func comparelines (line1, line2, headline,sep string,outfile *os.File ) {
        fields1 := strings.Split(line1, sep)
        fields2 := strings.Split(line2, sep)
        headers := strings.Split(headline, sep) 
        var minlen int = len(fields1) 
        if minlen >= len(fields2) {
            minlen = len(fields2) 
        } 
        for i := 0; i < minlen; i++ {
            if  len(fields1) < i && len(fields2) < i {
                break
            }  
            //fmt.Println("fields1 ", fields1[i], "fields2 ", fields2[i]  ) 
            if fields1[i] != fields2[i] {
                recid :=fields1[0]
                recname := fields1[1] 
                diff := recid +sep+recname+ sep+ headers[i] + sep+ fields1[i]+sep+fields2[i]+sep+"\n"  
                //fmt.Println(diff ) 
                outfile.Write([]byte(diff))
                diff =""
            }    
        } 
}    
func compare_csvs (csvfile1, csvfile2,sep string) ( outfilename string ){
    file1, err := os.Open(csvfile1)
    if err != nil {
        panic(err) 
    }
    defer file1.Close() 
    file2, err := os.Open(csvfile2)
    if err != nil {
        panic(err) 
    } 
    defer file2.Close()
    outfilename =strings.Replace(csvfile1, ".csv", "", 1) + "_"+strings.Replace(csvfile2, ".csv", "", 1) + ".csv"   
    _, err = os.Stat(outfilename)
    if err == nil {
        os.Remove(outfilename) 
    } 
    outfile, err := os.OpenFile(outfilename, os.O_CREATE|os.O_RDWR, 0666) 
    if err != nil {
        panic(err) 
    } 
    defer outfile.Close() 

    if err != nil {
        panic(err) 
    } 
    var lines1 map[string]string 
    lines1 = make(map[string] string)
    var lines2 map[string]string
    lines2 = make(map[string]string ) 
    r := bufio.NewReader(file1) 
    line, err := r.ReadString('\n')
    count := 1
    for err == nil {
        //fmt.Println(line)
        fields := strings.Split(line, sep) 
        if len(fields) < 3 {
            line, err = r.ReadString('\n') 
            continue
        }  
        recid := fields[0]
        _, ok := lines1[recid]
        num :=0
        for ok == true {
            temp := recid + ":dinakar:"+strconv.Itoa(num)
            num++
            _,ok = lines1[temp] 
            if ok == false { 
                recid = temp
                break
            } 
        }  
        lines1[recid] = line 
        line, err = r.ReadString('\n') 
        count++
    }
    if err != io.EOF {
        fmt.Println(err)

    }
    r = bufio.NewReader(file2) 
    line, err = r.ReadString('\n')
    count = 1
    for err == nil {
        //fmt.Println(line)
        fields := strings.Split(line, "|") 
        if len(fields) < 3 {
            line, err = r.ReadString('\n') 
            continue
        }  
        recid := fields[0]
        _, ok := lines2[recid]
        num :=0
        for ok == true {
            temp := recid + ":dinakar:"+strconv.Itoa(num) 
            num++
            _,ok = lines2[temp] 
            if ok == false {
                recid = temp
                break
            } 
        }  
        lines2[recid] = line 
        line, err = r.ReadString('\n') 
        count++
    } 
    if err != io.EOF {
        fmt.Println(err)

    }

    head := "Record ID"+ sep+"Record Name"+ sep+"Property"+sep+csvfile1+sep+csvfile2+sep+"\n"
    outfile.Write([]byte(head)  )
    ids := []string{"RULE ID"}
    var okid bool
    var headerline string
    for _, v := range ids { 
        headerline, okid = lines1[v]
        
    }
    if okid == false {
            fmt.Println("Headers do not exist, make sure headers exists in xml file")
            return ""
    } 
    for key := range(lines1) {
        line1, ok := lines1[key]
        line2, ok := lines2[key] 
        if ok == true {
            comparelines(line1,line2, headerline,sep, outfile) 
            delete(lines2, key)  
        } 
    }  
    for key := range(lines2) {
        fields := strings.Split(key, ":")
        if len(fields) > 1 {
            key = fields[0] 
        } 
        key = key + " does not exists in file1\n" 
        outfile.Write([]byte(key)) 
    }  
        
    return outfilename
}  
func main() {
    var xmlfile1 string
    var xmlfile2 string
    if  len(os.Args) > 2 {
        xmlfile1 = os.Args[1]
        xmlfile2 = os.Args[2] 

    } else {
        fmt.Println("Usage: ",os.Args[0], " <xml_file1> <xml_file2 [field separator <optional>] " )
        fmt.Println("xml file name extension has to be .xml") 
        help() 
        return
    }
    sep := "|"
    if len(os.Args) > 3 {
        sep = os.Args[3] 
    }  
    if ok := strings.Contains(xmlfile1, ".xml"); ok == false {
        fmt.Println("xml file name should end with .xml")
        help() 
        return
    } 
    if  ok := strings.Contains(xmlfile2, ".xml"); ok == false {
        fmt.Println("xml file name should end with .xml")
        help() 
        return

    }  
    xml2csv(xmlfile1)
    xml2csv(xmlfile2)
    csvfile1 := strings.Replace(xmlfile1, ".xml", ".csv", 1)
    csvfile2 := strings.Replace(xmlfile2, ".xml", ".csv", 1) 
    output := compare_csvs(csvfile1, csvfile2,sep) 
    _, err := os.Stat(csvfile1)
    if err == nil { 
        os.Remove(csvfile1)
    }
    _,err = os.Stat(csvfile2)
    if err == nil {  
        os.Remove(csvfile2)
    }
    out, err := os.Open(output)
    if err != nil {
        panic(err) 
    } 
    defer out.Close() 
    r := bufio.NewReader(out) 
    line, err := r.ReadString('\n')
    var lines map[string]string
    lines = make(map[string]string ) 
    for err == nil {
        lines[line] = ""
        line, err = r.ReadString('\n') 
    }
    file, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0666) 
    defer file.Close() 
    if err != nil {
        panic(err) 
    } 
    key, ok :=lines["RULE ID"]  
    if ok {
        file.Write([]byte(key))
        delete(lines, key)  
    } 
    for key  := range lines {
       file.Write([]byte(key))
        
    } 
    fmt.Println("Please see ", output, " for details for the comparison") 
    //filename = "junk.xml_schema"
    

}
