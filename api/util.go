package api

import (
	"time"
)

// Format Year/Month/Day/Hour/Min/Sec
const layout = "2006/01/2/15/04/05"

func mountDirectoryPathFromTime(time time.Time) string {
	return time.Format(layout)
}
