package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLine() string {
    fmt.Println("Please enter some text") 
	buf := bufio.NewReader(os.Stdin)
	line, err := buf.ReadString('\n')

	if err != nil {
		panic(err)
	}

	line = strings.TrimRight(line, "\n")

	return line
}

func main() {
	//input := readLine() // <= Hello World
	//fmt.Println(input)  // => Hello World
    file := "relb_aip.csv"
    name := strings.Split(file, ".")[0]  
    fmt.Println(name) 
    t :=[]int{0,1,2,3,4}  
    fmt.Println(t) 
    t = t[1:] 
    fmt.Println(t) 
}

