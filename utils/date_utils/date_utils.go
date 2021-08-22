package date_utils

import "time"

const (
	TimeStamp  = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(TimeStamp)
}
