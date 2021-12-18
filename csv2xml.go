package main

import ("fmt"
        "encoding/xml"
        "strings"
        "os"
        "io"
        "encoding/csv"
        "path/filepath"
)
var help string = `========================
csv2xml program will take csv file (default separator is comma (,)) and converts to xml file with same name with xml extension. Structure of xml file is predefined or agreed upon.   
usage: 
csv2xml <file_name.csv> 
will produce a file called file_name.xml
example:
csv2xml testing.csv  -->output testing.xml 
csv2xml "c:\user\dsdesai\my documents\testing.csv" 

===========================
Author: Dinakar Desai, Ph.D.
Date: Nov 9, 2017

If you experience any problems with the program, please contact above author.
`
const (
    Header = `<PCDPdataset xsi:schemaLocaiton="http://www.pecarn.org/pcdp
    file://Z:/pcdp/trunk/WIKI/pcdp8.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://www.pecarn.org/pcdp">` + "\n<Site>TCBC</Site>\n"
) 
/*
type Record struct {
    XMLName xml.Name `xml:"Site"`
    PcdpRecs []PCDPrecord `xml:"PCDPrecords"`
} 
*/
type PCDPrecord struct {
    XMLName xml.Name `xml:"PCDPrecord"`
    MedRecNbr string `xml:MedRecNbr"`
    BirthDate string `xml:"BirthDate"`
    Gender string `xml:"Gender"`
    Race string `xml:"Race"`
    Ethnicity string `xml:"Ethnicity"`
    Zip string `xml:"Zip"`
    TriageCat string `xml:"TriageCategory"`
    ChiefComp []string `xml:"CheifComplaint,omitempty"`
    ProcCode []string `xml:"ProcedureCode,omitempty"`
    Icd10ProcCode []string `xml:"ICD10ProcedureCode,omitempty"`
    DxCode []string `xml:"DXCode,omitempty"`
    Icd10DxCode []string `xml:"ICD10DxCode,omitempty"`
    ECode []string `xml:"ECode,omitempty"`
    Payer string `xml:"Payer,omitempty,omitempty"`
    EdDisposition string `xml:"EdDisposition,omitempty"`
    TriageDate string `xml:"TriageDate,omitempty"`
    TriageTime string `xml:"TriageTime,omitempty"`
    DischargeDate string `xml:"DischargeDate,omitempty"`
    DischargeTime string `xml:"DischargeTime,omitempty"`
    ModeOfArrival string `xml:"ModeOfArrival,omitempty"`
} 

func deleteBlank(in []string ) []string {
    temp := make([]string, 0)  
    for _,value := range in {
        if value != "" {
            temp = append(temp, value)  
        }  
    }  
    return temp
} 
func csv2xml(inFile, outFile string) {
    var lineCount int = 0;
    csvFile, err := os.Open(inFile)
    if err != nil {
        fmt.Println("Error:", err)
		return
    } 
    defer csvFile.Close()
    f, err := os.Create(outFile) 
    if err != nil {
        panic(err)  
    } 
    defer f.Close() 
    reader := csv.NewReader(csvFile)
   // var lineCount int= 0
 //   var s Site
    PcdpRecs := make([]PCDPrecord, 0 )
    for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
            lineCount += 1
        
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
        if (lineCount == 1) { continue } 
             
        var rec PCDPrecord
        rec.MedRecNbr= record[1] 
        rec.BirthDate = record[2] 
        rec.Gender = record[3]
        rec.Race = record[4] 
        rec.Ethnicity = record[5] 
        rec.Zip = record[6] 
        rec.TriageCat = record[7] 
        rec.ChiefComp = record[8:11] 
        rec.ProcCode = deleteBlank(record[11:26])
        rec.Icd10ProcCode = deleteBlank(record[26:41])
        rec.DxCode = deleteBlank(record[41:56])
        rec.Icd10DxCode = deleteBlank(record[56:71])
        rec.ECode = deleteBlank(record[71:76])
        rec.Payer = record[76] 
        rec.EdDisposition = record[77] 
        rec.TriageDate = record[78] 
        rec.TriageTime = record[79] 
        rec.DischargeDate = record[80] 
        rec.DischargeTime = record[81] 
        rec.ModeOfArrival = record[82] 
        PcdpRecs = append (PcdpRecs, rec) 
    
    }
    if xmlstring, err := xml.MarshalIndent(PcdpRecs,"","   "); err == nil {
        xmlstring = []byte(xml.Header +Header+ string(xmlstring) )
        //fmt.Printf("%s\n", xmlstring) 
        fmt.Fprintf(f,string(xmlstring)) 
        fmt.Fprintf(f, "\n</PCDPdataset>") 
    }

} 
func main() {
    var inFile, outFile string
    if len(os.Args) == 1 {
        fmt.Println(help) 
        return
    } else {
        inFile = os.Args[1]  
    }  
    if _, err := os.Stat(inFile); err == nil { 
        _,basename := filepath.Split(inFile)  
        basename = strings.Replace(strings.ToLower(basename), "csv", "xml",1) 
        outFile = basename
        csv2xml(inFile, outFile) 
    } else {
        fmt.Println(inFile, " does not exist, please provide correct path to the existing file\n\n", help) 
    }  
}  


