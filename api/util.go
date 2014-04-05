package api

import (
    "fmt"
	"time"
	"os"
)

// Format Year/Month/Day/Hour/Min/Sec
const layout = "2006/01/2/15/04/05/"

func mountDirectoryPathFromTime(time time.Time) string {
	return time.Format(layout)
}


func CreateDir(path string) error {
    return os.MkdirAll(path, os.ModeDir)
}

func check(err error){
    if err != nil{
        fmt.Println(err)
    }
}
