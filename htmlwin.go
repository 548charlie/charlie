package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os/exec"
    "time"
)

const (
    port = 8765
)

var guestList []string

// indexHandler serves the main page
func indexHandler(w http.ResponseWriter, req *http.Request) {
    t := template.New("index.html")
    t, err := t.Parse(indexHTML)
    if err != nil {
        message := fmt.Sprintf("bad template: %s", err)
        http.Error(w, message, http.StatusInternalServerError)
    }

    t.Execute(w, guestList)

}

// openBrowser waits one second and then open web browser on us
func openBrowser() {
    time.Sleep(time.Second)
    url := fmt.Sprintf("http://localhost:%d", port)
    exec.Command("start", url).Start()
}

// addHandler add a name to the names list
func addHandler(w http.ResponseWriter, req *http.Request) {
    guest := req.FormValue("name")
    if len(guest) > 0 {
        guestList = append(guestList, guest)
    }

    http.Redirect(w, req, "/", http.StatusFound)
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/add", addHandler)
    go openBrowser()
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}


var indexHTML = `
<!DOCTYPE html>
<html>
    <head>
        <title>Guest Book ::Web GUI</title>
    </head>
    <body>
        <h1>Guest Book :: Web GUI</h1>
        <form action="/add" method="post">
        Name: <input name="name" /><submit value="Sign Guest Book">
        </form>
        <hr />
        <h4>Previous Guests</h4>
        <ul>
            {{range .}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    </body>
</html>
`

