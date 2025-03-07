package tools

import "time"

func TaskID() string {
	return time.Now().Format("20060102150405")
}
