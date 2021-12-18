package main

import (
    "fmt"
    "os"        
    "io"
    "bufio"
    "regexp"
    "strings"
    "strconv"
    "encoding/csv"
    "time"
    "path/filepath" 
) 
var help string = `========================
colstat_index program will take csv files (default separator is comma (,)) and create an index file that will be consumed by onbase
file_type is either image or visit
usage: 
colstat_index <file_type> <config_filename>
will produce a index.txt, index1.txt etc based on number of lines 
example:
colstat_index <file_type> <filename.conf> (if present in same directory as program) 
colstat_index image "c:\user\dsdesai\my documents\filename.conf" 
colstat_index visit "c:\user\dsdesai\my documents\filename.conf" 

structure of conf file is as follows:
image_file = "<full path to image.csv file>
demograph_file = <full path to demographic.cvs
provider_file = <full path to provider file>
visits_file = <full path to visits file>
image_doc_path_prefix = <unc path to prefix to image files>
visit_doc_path_prefix = <unc path to prefix to visit doc files>
lines_per_file = 100 (-1 will create one file)  is number of lines per index file and index files will be suffixed with 1 and increments

example of conf file:
image_file = c:/tch/colstat_dat/images.csv
demograph_file = c:/tch/colstat_dat/demographic.csv
provider_file = c:/tch/colstat_dat/provider_info.csv
visits_file = c:/tch/colstat_dat/visits.csv
image_doc_path_prefix  = c:/tch/test/images
visit_doc_path_prefix = c:/tch/test/visits
lines_per_file = 100

===========================
Author: Dinakar Desai, Ph.D.
Date: Nov 9, 2017

If you experience any problems with the program, please contact above author.
`
var image_file, demograph_file,provider_file,visits_file, image_doc_path_prefix, visit_doc_path_prefix string
var lines_per_file int
func removeSplChars (data string)(out string) {
    src := strings.Replace(data, "\\", "", -1)
    src = strings.Replace(src, "\\n", "",-1) 
    re := regexp.MustCompile("[|^~&]" )
    out = re.ReplaceAllString(src, "") 
    return out
}   
func getDateFormat(dt, format string) ( fmtStr string) {
    var ss string
    if format == "MMDDYYYY" {
        if dt == "" {dt = "01/01/2000"} 
        dst := strings.Split(dt, "/")
        month, day,year := dst[0],dst[1],dst[2]
        dt =fmt.Sprintf("%02s/%02s/%04s", month,day,year) 
       t, err := time.Parse("01/02/2006", dt)
        if err != nil {  
            fmt.Println("MMDDYYYY--", err)         
        }
        //s := t.Format("20060102150405")
        ss = t.Format("01022006") 
        
    } else if (format == "YToS"){
        if dt == "" {dt = "2000/01/01"} 
        dt = strings.Replace(dt, "-", "/", -1) 
        dst := strings.Split(dt, "/")
        year, month,day := dst[0],dst[1],dst[2]
        dt =fmt.Sprintf("%04s/%02s/%02s", year,month,day)
        t, err := time.Parse("2006/01/02", dt)
        if err != nil {fmt.Println("YToS--", err) }
        ss = t.Format("20060102150405") 
    }   
    return ss 
}  
func readRows(filename string, col_id int) (rows map[string][]string  )  {
    rows =make(map[string][]string  ) 
    //fmt.Println("working on:", filename) 
    file, err := os.Open(filename)
    defer file.Close()
    if err != nil {
        fmt.Println(filename, " ", err)
        return
    } else {
        file.Seek(0, os.SEEK_SET)
        r := csv.NewReader(file)
        r.TrailingComma = true
        r.TrimLeadingSpace = true
        row , err := r.Read()
        if err != nil {
            return rows
        } 
        for err == nil {
            if err == io.EOF {
                break
            } 
            if row[col_id] != "" {
                rows[row[col_id] ] = row 
            }  
            row,err = r.Read() 
        } 
    } 
   return rows 
}
func get_viewer(ext string) string {
    exts := []string{"tiff","bmp","jpg","tif","jpeg","png"    }
    ext = strings.ToLower(ext) 
    contain := strings.Contains(ext, "pdf")
    var viewer string = "not_defined" 
    if contain == true {
        viewer =  "16" 
    }  else {
        for _, ex := range exts {
            contain = strings.Contains(ext, ex) 
            if contain == true {
                viewer= "2" 
            } 
        } 
    }
    return viewer
} 
func create_index (file_type string){
    //img_data := readRows(image_file, 5) 
    demo_data :=readRows(demograph_file,0) 
    provider_data := readRows(provider_file,0) 
    visits_data := readRows(visits_file, 0) 
    file_name := "index_image.txt"
    f, err := os.OpenFile(file_name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) 
    defer f.Close() 
    if err != nil { 
        panic(err)         
    } 
/*
     for i, value := range img_data {
        if i == "Patient ID" { 
             fmt.Println("img_data:", i, "->", value)
        }
    }
    for i, value := range demo_data {
        if i == "Patient ID" { 
            fmt.Println("demo_data:", i, "->", value) 
        }
    } 
    for i, value := range provider_data {
        if i == "PCC Provider Name"{ 
            fmt.Println("prov_data:",i, "->", value)
        }
    } 
    for i, value := range visits_data {
        if i == "Encounter ID" { 
            fmt.Println("visit_data:",i, "->", value)
        }
    }
    fmt.Println("mrn|csn|pat_Lastname|pat_firstName|Gender|DOB|Provider_ID|Provider_LastName|Provider_firstName|doc_type|epic_doc_comment|service_date|doc_sign_date|image_filename") 
    */
    var mrn,csn, pat_ln, pat_fn, gender,dob,pr_id,pr_ln,pr_fn,doc_type,epic_doc_cmt, service_dt,doc_sign_dt,img_path,img_fn, ext,image_fn, viewer string

    if file_type == "image"  { 
        file, err := os.Open(image_file)
        defer file.Close()
        if err != nil {
            fmt.Println(image_file, " ", err)
            return
        } else {
            file.Seek(0, os.SEEK_SET)
            r := csv.NewReader(file)
            r.TrailingComma = true
            r.TrimLeadingSpace = true
            row , err := r.Read()
            if err != nil {
                fmt.Println(err) 
            } 
            for err == nil {
                if err == io.EOF {
                    break
                } 
            mrn = row[5] 
            if mrn == "Patient ID" {
                row,err = r.Read() 
                continue
            } 
            if _,ok := demo_data[mrn]; ok {   
                csn = row[3]
                pat_ln = demo_data[mrn][1]
                pat_fn = demo_data[mrn][2]
                dob = demo_data[mrn][6]  
                dob = getDateFormat(dob, "MMDDYYYY") 
                gender = string(demo_data[mrn][5][0])   
                if csn == "" {
                    pr_id = provider_data["DEFAULT"][1]
                    pr_name := strings.Split(provider_data["DEFAULT"][2],",")
                    pr_ln, pr_fn = pr_name[0],pr_name[1]  
                } else if _, ok := visits_data[csn]; ok {
                    pr_name := visits_data[csn][4]
                    if _, ok := provider_data[pr_name]; ok {  
                        tch_pr_name := provider_data[pr_name]
                        pr_id =tch_pr_name[1]
                        tpr_name := strings.Split(tch_pr_name[2], "," )
                        pr_ln, pr_fn = tpr_name[0],tpr_name[1]  
                    } else {
                        pr_id = provider_data["DEFAULT"][1]
                        pr_name := strings.Split(provider_data["DEFAULT"][2],",")
                        pr_ln, pr_fn = pr_name[0],pr_name[1]  
                    } 
                } 
                epic_doc_cmt = row[6]
                doc_type = "External Historic Document"
                if epic_doc_cmt == "" {
                    epic_doc_cmt = row[7]
                    if epic_doc_cmt == "" {
                        epic_doc_cmt = "External Document"
                        doc_type = "Exernal Historic Document"
                    } 
                } 
                epic_doc_cmt = removeSplChars(epic_doc_cmt) 
                service_dt = row[4]
                img_path = row[8]  
                img_fn = row[9]  
                ext = filepath.Ext(img_fn) 
                viewer = get_viewer(ext)  
                if service_dt == "" {
                    var sr_dt_match string =`(\d{4,}/\d{2,}/\d{2,})` 
                    reg,err := regexp.Compile(sr_dt_match)
                    if err != nil {
                        fmt.Println(err) 
                    } 
                    matched := reg.Find([]byte(img_path))
                    if len(matched) > 0 {
                        service_dt =string([]byte(matched)  ) 
                    }  
                }
                dst := strings.Split(service_dt,"/" ) 
                if len(dst) == 3 { 
                    yr, mt, da := dst[0],dst[1],dst[2]  
                    if len(yr)  == 4{ 
                        service_dt = yr +"/" +mt+"/" +da
                    }else {
                        service_dt = da+"/"+ yr+ "/" +mt
                    }
                }
                service_dt = getDateFormat(service_dt,"YToS") 
                var year, month string
                year = service_dt[0:4]
                month = service_dt[4:6]
                day,_ := strconv.ParseInt( service_dt[6:8], 0,64)
                if year == "2019" && month  == "01" && day > 4 {
                    row,err = r.Read() 
                    continue
                } 
                doc_sign_dt = service_dt
                image_fn = image_doc_path_prefix +strings.TrimSpace(img_path)+"\\"+strings.TrimSpace(img_fn)
                image_fn = strings.Replace(image_fn, "/", "\\", -1) 
                var items []string
                items = append(items, mrn,csn,pat_ln,pat_fn,dob,gender,pr_id,pr_ln,pr_fn,doc_type,epic_doc_cmt,service_dt,doc_sign_dt,viewer,image_fn) 
                line := strings.Join(items[:], "|" )  
            //    fmt.Println(line) 
                    
                fmt.Fprintf(f,"%s\n",line)
                
                }  
                //fmt.Println(mrn, "csn:",csn,"pat_ln:", pat_ln, "pat_fn:", pat_fn, "gender:",gender,"dob:", dob, "pr_id:", pr_id, "prn_ln:", pr_ln, "pr_fn:", pr_fn, "doc_type:", doc_type,"epic_doc_cmt:",epic_doc_cmt,"service_dt:", service_dt, "doc_sign_dt:", doc_sign_dt, "filename:", image_fn)


                row,err = r.Read() 
            }
        }
    } else if (file_type == "visit") {
        file_name = "index_visit.txt" 
        f, err := os.OpenFile(file_name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) 
       defer f.Close() 
       if err != nil { 
            panic(err)         
       } 
        for csn = range visits_data {
            if csn == "Encounter ID" {continue} 
            mrn = visits_data[csn][1]
            if _,ok := demo_data[mrn]; ok {   
                pat_ln = demo_data[mrn][1]
                pat_fn = demo_data[mrn][2]
                dob = demo_data[mrn][6]  
                dob = getDateFormat(dob,"MMDDYYYY") 
                gender = string(demo_data[mrn][5][0])  
                if csn == "" {
                    pr_id = provider_data["DEFAULT"][1]
                    pr_name := strings.Split(provider_data["DEFAULT"][2],",")
                    pr_ln, pr_fn = pr_name[0],pr_name[1]  
                } else if _, ok := visits_data[csn]; ok {
                    pr_name := visits_data[csn][4]
                    if _, ok := provider_data[pr_name]; ok {  
                        tch_pr_name := provider_data[pr_name]
                        pr_id =tch_pr_name[1]
                        tpr_name := strings.Split(tch_pr_name[2], "," )
                        pr_ln, pr_fn = tpr_name[0],tpr_name[1]  
                    } else {
                        pr_id = provider_data["DEFAULT"][1]
                        pr_name := strings.Split(provider_data["DEFAULT"][2],",")
                        pr_ln, pr_fn = pr_name[0],pr_name[1]  
                    } 
                } 
                doc_type = "External Historic Documents"
                epic_doc_cmt = "Telephone Encounter"
                service_dt = visits_data[csn][2]  
                service_dt = strings.Replace(service_dt,"-","", -1) 
                service_dt = strings.Replace(service_dt,":", "", -1)
                service_dt = strings.Replace(service_dt," ", "", -1) 
                service_dt = service_dt[0:14] 
                img_fn = visits_data[csn][6]   
                ext = filepath.Ext(img_fn) 
                viewer = get_viewer(ext)  

                /*
                if service_dt == "" {
                    var sr_dt_match string =`(\d{4,}/\d{2,}/\d{2,})` 
                    reg,err := regexp.Compile(sr_dt_match)
                    if err != nil {
                        fmt.Println(err) 
                    } 
                    matched := reg.Find([]byte(img_path))
                    if len(matched) > 0 {
                        service_dt =string([]byte(matched)  ) 
                    }  
                } 
                */
                var year, month string 
                year = service_dt[0:4]
                month = service_dt[4:6]
                day,_ := strconv.ParseInt( service_dt[6:8], 0,64)
                if year == "2019" && month  == "01" && day > 4 {continue} 
                doc_sign_dt = service_dt
                image_fn := visit_doc_path_prefix +"/"+strings.TrimSpace(img_fn)  

                image_fn = strings.Replace(image_fn, "/", "\\", -1)
                image_fn = strings.Replace(image_fn, ".html", ".pdf", 1) 
                var items []string
                items = append(items, mrn,csn,pat_ln,pat_fn,dob,gender,pr_id,pr_ln,pr_fn,doc_type,epic_doc_cmt,service_dt,doc_sign_dt,viewer,image_fn) 
                line := strings.Join(items[:], "|" )  
                //fmt.Println(line) 
                fmt.Fprintf(f,"%s\n",line)
                        
                  
               // fmt.Println("mrn:", mrn, "csn:",csn,"pat_ln:", pat_ln, "pat_fn:", pat_fn, "gender:",gender,"dob:", dob, "pr_id:", pr_id, "prn_ln:", pr_ln, "pr_fn:", pr_fn, "doc_type:", doc_type,"epic_doc_cmt:",epic_doc_cmt,"service_dt:", service_dt, "doc_sign_dt:", doc_sign_dt, "filename:", image_fn)

            } 
        } 
    }  
}


func main() {
    var config,file_type string
    if (len(os.Args) > 2 ) {
        file_type = os.Args[1] 
        config = os.Args[2] 
    } else {
       fmt.Println(help)  
       return
    } 
    if file_type != "image" && file_type != "visit" {return} 
    cfile, err := os.Open(config)
    defer cfile.Close()
    if err != nil {
        panic(err) 
    } 
    cf := bufio.NewScanner(cfile)
    for cf.Scan(){
        line:= cf.Text()
        if line == "" {continue}
        words := strings.Split(line, "=")
        name, value := strings.TrimSpace(words[0]),strings.TrimSpace(words[1])       
        if name == "image_file" {
            image_file = value 
        } else if name== "demograph_file" {
            demograph_file = value 
        } else if name == "visits_file" {
            visits_file = value 
        } else if name  == "provider_file" {
            provider_file = value 
        } else if name == "image_doc_file_path" {
            image_doc_path_prefix = value
        } else if name == "visit_doc_file_path" {
            visit_doc_path_prefix = value
        } else if name == "lines_per_file" {
            lines_per_file,_= strconv.Atoi(value)
        } 
    } 


    create_index(file_type) 

}  
