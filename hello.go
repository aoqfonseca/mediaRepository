package main

import (
    "log"
    "net/http"
    "runtime"
    "github.com/gorilla/pat"
)

func main() {

    r := pat.New()
    
    //Creanting routes
    
    r.Get("/healthcheck", http.HandlerFunc(HealthCheckHandler))
    r.Get("/upload", http.HandlerFunc(PostFormUploadFile))
    r.Post("/upload", http.HandlerFunc(UploadFileHandler))
    
    log.Println("Iniciando sistema o")
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    log.Println("Servidor escutando na porta 43211")
    http.ListenAndServe(":4321", r)
}


