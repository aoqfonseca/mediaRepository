package main

import (
    "log"
    "net/http"
    "runtime"
    "github.com/gorilla/pat"
    "github.com/mediaRepository/api"
)

func main() {

    r := pat.New()
    
    //Creating routes
    r.Get("/healthcheck", http.HandlerFunc(api.HealthCheckHandler))
    r.Get("/upload", http.HandlerFunc(api.PostFormUploadFile))
    r.Post("/upload", http.HandlerFunc(api.UploadFileHandler))
    
    log.Println("Staring server: getting num cpu")
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    log.Println("Server listening on 4321")
    http.ListenAndServe(":4321", r)
}


