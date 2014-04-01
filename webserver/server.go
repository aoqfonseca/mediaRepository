package main

import (
    "flag"
	"github.com/gorilla/pat"
    "github.com/tsuru/config"
	"github.com/mediaRepository/api"
	"log"
	"net/http"
	"runtime"
)

func main() {

    configFile := flag.String("config", "/etc/config.yml", "Media Repository Path Config file")
    port        := flag.String("port", ":4321", "Listening por. Example: :4321")
    flag.Parse()


    err := config.ReadAndWatchConfigFile(*configFile)
    if err !=nil{
        fmt.Println("Could not find config file")
    }

	r := pat.New()

	//Creating routes
	r.Get("/healthcheck", http.HandlerFunc(api.HealthCheckHandler))
	r.Get("/upload", http.HandlerFunc(api.PostFormUploadFile))
	r.Post("/upload", http.HandlerFunc(api.UploadFileHandler))

	log.Println("Staring server: getting num cpu")
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println("Server listening on 4321")
	http.ListenAndServe(port, r)

}
