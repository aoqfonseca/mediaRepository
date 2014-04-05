package api

import (
	"fmt"
    "github.com/tsuru/config"
	"io/ioutil"
    "net/http"
    "time"
)

const (
	IMAGE_TYPES       = "image/(gif|p?jpeg|(x-)?png)"
	VIDEO_TYPES       = "(video|realmedia)"
)

func UploadFileHandler(w http.ResponseWriter, req *http.Request) {
    var pathToSave, err = config.GetString("photo_storage_path")
    now := time.Now()
    var directory = mountDirectoryPathFromTime(now)
    check(err)

	file, handler, err := req.FormFile("file")
	check(err)
	
	data, err := ioutil.ReadAll(file)
	check(err)
	
    directory = pathToSave + directory
	err = CreateDir(directory)
	check(err)
	
	var path_file = directory +  handler.Filename
	err = ioutil.WriteFile(path_file, data, 0777)
	check(err)
	
	fmt.Fprintf(w, "SUCCESS")
}

func PostFormUploadFile(w http.ResponseWriter, req *http.Request) {
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
