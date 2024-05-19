package helper

import "time"

func ParseOnlyDate(date string) string {
	dateParse, _ := time.Parse(time.RFC3339, date)

	return dateParse.Format("2006-01-02")
}

func ParseTimeOnly(t string) string {
	dateParse, _ := time.Parse(time.RFC3339, t)
	return dateParse.Format("15:04:05")
}
