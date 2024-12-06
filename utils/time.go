package utils

import "time"

const (
	shortLayout = "2006-01-02"
	longLayout  = "2006-01-02 15:04:05"
)

func ParseTime(timeStr string) time.Time {
	t, _ := time.Parse(shortLayout, timeStr)
	return t
}
