package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "strings"
    "sort"

) 
junk
var help string = `========================
diffmapper program will compare two mapper files and displays the
difference between the files.

usage:
diffmapper <root of repository> <env1> <env2> <file_pattern|filename without extension>

root of repository is directory path till including local_storage 
eg: C:\tch\rhapsody\local_storage\editst.gslb.ad.texaschildrenshospital.org-3041\Definitions\cac_adt.mdf (2b9ed47) is my repository for cac_adt mdf. So the required path is:  C:\tch\rhapsody\local_storage

if there are spaces in the path, please quote the path as shown in the example below.

example:
diffmapper "c:\tch\rhapsody\local_storage" dev test alaris_adt
diffmapper "c:\user\dsdesai\my documents\local_definition_storage" dev test oru
diffmapper "c:\tch\rhapsody\local_storage" dev test adt

===========================
Author: Dinakar Desai, Ph.D.
Date: Nov 9, 2017

If you experience any problems with the program, please contact above author.
`
func readFile( filename string )map[string]int {
    fh, err := os.Open(filename)
    fields := make(map[string]int ) 
    if err !=  nil {
        fmt.Println(err) 
    } 
    defer fh.Close()
    fs := bufio.NewScanner(fh)
    var linecount,tCurle, mCurle, typeStart, mapStart int = 0,0,0,0,0
    var aType, amap string
    curOpen, _ := regexp.Compile("{")
    curClose,_ := regexp.Compile("}") 

    mapReg, _ :=regexp.Compile("^map ")
    typeReg, _ :=regexp.Compile("^type")  
    for fs.Scan() {
        line := fs.Text()
        linecount += 1
        tline := strings.TrimSpace(line)
        match := typeReg.Find([]byte(tline)) 
        if len(match) > 0 {
           typeline := strings.Split(tline, " ")
           aType = typeline[1] 
           typeStart = 1
        }  
        match = curOpen.Find([]byte(tline))
        if len(match) > 0 {
            tCurle += 1
        } 
        match = curClose.Find([]byte(tline))
        if len(match) > 0 {
            tCurle -=1
        }  
       this is testing 
        if tCurle >= 0 && typeStart == 1{
            fields["type|"+aType+ "|" + tline] = linecount 

        }
        if tCurle == 0 {
            typeStart = 0
        } 
        match = mapReg.Find([]byte(tline)) 
        if len(match) > 0 {
           typeline := strings.Split(tline, " ")
           amap = typeline[1] 
           mapStart = 1
        }  
        match = curOpen.Find([]byte(tline))
        if len(match) > 0 {
            mCurle += 1
        }
        match = curClose.Find([]byte(tline))

        if len(match) > 0 {
            mCurle -= 1
        }  
        
        if mCurle >= 0 && mapStart == 1 {
            fields["map|" +amap+ "|" + tline] = linecount 
        }
        if (mCurle == 0 || tCurle == 0 ) {
            if _,ok := fields[tline]; ! ok {
                fields[tline] = linecount 
            } else {
                //fmt.Println("line does exist in map " + tline) 
            }  
        }  

    }  

   // for key,value := range fields {
     //   fmt.Printf("key %s value %d\n", key, value) 
    //} 
    return fields

}  

func compareMaps (lines1, lines2 map[string]int, env1, env2 string ) {
    diffs1 := make(map[int]string)
    diffs2 := make(map[int]string ) 
    display1 := make(map[int]string)
    display2 := make(map[int]string ) 


    for key, value := range lines1 {
        display1[value] = key 
        if _, ok := lines2[key]; ! ok {
            diffs1[value] = key 
        } 
         
    } 
    for key, value := range lines2 {
        display2[value] = key 
        if _, ok := lines1[key]; ! ok {
            diffs2[value] = key 
        }  
    } 
    /*
    for key, value := range diffs1 {
        fmt.Printf("%d --%s\n", key,value) 

    } 
    for key, value := range diffs2 {
        fmt.Printf("%d -- %s\n", key,value) 
    } 
    */
    keymap := make(map[int]int) 
    var keys []int
    for k :=range diffs1 {
        keys = append(keys, k)
        keymap[k] =k
    } 
    for k := range diffs2 {
        if _, ok := keymap[k]; ! ok {  
            keys = append(keys, k) 
        }
    } 

    sort.Ints(keys)
    if len(keys) > 0 {  
        fmt.Printf("Line Number|%35s|%40s\n", env1, env2) 
        for _,k := range keys {
                fmt.Printf("%d|%35s|%40s\n", k, diffs1[k], diffs2[k]) 
        } 
    }else {
        fmt.Println("No differences found between environments") 
    } 

} 
func compareLocal(files1, files2 []string  ) {
    for i,file1 := range files1 {
        file2 := files2[i] 
        lines1 := readFile(file1)
        lines2 := readFile(file2)
        compareMaps(lines1, lines2, file1,file2) 
    } 
}  
func main() {
    var fileRoot, filePath1,filePath2, env1, env2, filePattern string
    var server1, server2 string
    if (len(os.Args) == 3 ) {
        file1pat := os.Args[1]
        file2pat := os.Args[2] 
        files1,_ := filepath.Glob(file1pat)
        files2,_ := filepath.Glob(file2pat)
        if (len(files1) == 0 ) {
            fmt.Printf("no files exist with %s name", file1pat) 
        } 
        if (len(files2) == 0  ) {
            fmt.Printf("no files exist with %s name", file2pat) 
        }  
        compareLocal(files1, files2) 
        return
            
    }  
    if (len(os.Args) < 5  ) {
        fmt.Println(help) 
        return  
    }   
    fileRoot = os.Args[1]

    env1 = strings.ToLower(os.Args[2])
    env2 = strings.ToLower(os.Args[3])
    if env1 == "dev" {
        server1 = "editst.gslb.ad.texaschildrenshospital.org-3042"
    } else if (env1 == "test" ) {
        server1 =  "editst.gslb.ad.texaschildrenshospital.org-3041"
    } else if (env1 == "prod") {
        server1 =  "ediprd.gslb.ad.texaschildrenshospital.org-3041"
    } 

    if env2 == "dev" {
        server2 = "editst.gslb.ad.texaschildrenshospital.org-3042"
    } else if (env2 == "test" ) {
        server2 =  "editst.gslb.ad.texaschildrenshospital.org-3041"
    } else if (env2 == "prod") {
        server2 =  "ediprd.gslb.ad.texaschildrenshospital.org-3041"
    } 
    filePattern = "*" +  os.Args[4] + "*"
    sep := string(os.PathSeparator)
    filePath1 = fileRoot+sep + server1 + sep + "Definitions"+sep +filePattern +sep + filePattern +".txt"
    filePath2 = fileRoot+sep + server2 + sep + "Definitions"+sep +filePattern +sep 
    files1,_ := filepath.Glob(filePath1 ) 
    for _, file := range files1 {
        _, basename := filepath.Split(file) 
        files2,_ := filepath.Glob(filePath2 + basename)  
        if len(files2) > 0 {  
            filename := files2[0] 
            _, err := os.Stat(filename )
            if err == nil {
                lines1 := readFile(file)
                lines2 := readFile(filename) 
                fmt.Println("Comparing " + basename ) 
                compareMaps(lines1, lines2, env1, env2)          
            } else {
                fmt.Println(filename  + " does not exist") 
            } 
        } else {
            fmt.Println("Corresponding file " + basename  + " does not exist in " + env2) 
        } 
    }  
}  

