package api

import (
    "fmt"
	"time"
	"os"
)

// Format Year/Month/Day/Hour/Min/Sec
const layout = "2006/01/02/15/04/"

func mountDirectoryPathFromTime(time time.Time) string {
	return time.Format(layout)
}


func CreateDir(path string) error {
    return os.MkdirAll(path, 0777)
}

func check(err error){
    if err != nil{
        fmt.Println(err)
    }
}
