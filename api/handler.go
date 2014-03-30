package api 

import (
    "log"
    "fmt"
    "net/http"
    "io/ioutil"
)

func UploadFileHandler(w http.ResponseWriter, req *http.Request) {
    log.Println
    file, handler, err := req.FormFile("file")
    if err != nil {
        fmt.Println(err)
    }
    data, err := ioutil.ReadAll(file)
    if err != nil {
            fmt.Println(err)
    }
    err = ioutil.WriteFile(handler.Filename, data, 0777)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(w, "SUCCESS")
}

func PostFormUploadFile(w http.ResponseWriter, req *http.Request){
    fmt.Fprintf(w, "<html><body>")
    fmt.Fprintf(w, "<form enctype=\"multipart/form-data\" action=\"http://localhost:4321/upload\" method=\"post\">")
    fmt.Fprintf(w, "  <input type=\"file\" name=\"file\" />")
    fmt.Fprintf(w, "  <input type=\"submit\" value=\"upload\" />")
    fmt.Fprintf(w, "</form>")
    fmt.Fprintf(w, "</body></html>")
}


func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "WORKING")
}
