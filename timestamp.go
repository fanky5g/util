package util

import "time"

// GetCurrentTimestamp formats current time in "2006-01-02 15:04:05" layout
func GetCurrentTimestamp() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
