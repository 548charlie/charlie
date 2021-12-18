package main
import (
        "net/http"
        "fmt"
        "strings"
        ) 

func helloHandler(w http.ResponseWriter, r *http.Request) {
    remPartOfUrl := r.URL.Path[len("/hello/"): ] 
    fmt.Fprintf(w, "Hello %s!!", remPartOfUrl) 
} 

func shouthelloHandler(w http.ResponseWriter, r *http.Request){
    remPartOfUrl := r.URL.Path[len("/shoutehllo/"): ] 
    fmt.Fprintf(w, "Hello %s!!!!", strings.ToUpper(remPartOfUrl)) 
}  
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Inside Handler") 
    fmt.Fprintf(w, "Hello world from my Go program!!!") 
} 

func main() {
    http.HandleFunc("/hello/", helloHandler) 
    http.HandleFunc("/shouthello/", shouthelloHandler) 
    http.ListenAndServe("localhost:9999", nil) 
} 
