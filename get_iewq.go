package main

import (
	"encoding/csv"
	"fmt"
	"os"
	//"strconv"
	"io"
	"strings"
)

func main() {
	var filename, delim string

    if len(os.Args)  > 1 {
        filename = os.Args[1] 
    } else {     
    fmt.Println("Please provide the name of csv file with cer jxport")
            return
    }  
	var start int = 0
	var finish int = 0

	wfile, err := os.Create("junk.csv")
	if err != nil {
		fmt.Println("could not create file")
		return
	}
	defer wfile.Close()

	writer := csv.NewWriter(wfile)
	defer writer.Flush()
	delim = ","
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file does not exist")
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ([]rune(delim))[0]
	reader.TrailingComma = true
	reader.TrimLeadingSpace = true
	row, err := reader.Read()
	if err != nil {
		return
	}
	for err == nil {
		if err == io.EOF {
			break
		}
		name := row[1]
		if strings.HasPrefix(name, "IEWQ") {
			start = 1
			finish = 0
		} else if start == 1 && name == "" {
			finish = 0
		} else {
			finish = 1
			start = 0
		}
		if finish == 0 {
			err = writer.Write(row)

			fmt.Println(row)
		}
		//fmt.Println(row)
		row, err = reader.Read()
	}
}
