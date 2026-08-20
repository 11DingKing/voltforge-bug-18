package protocol

import "time"

func ValidityWindowDay(t time.Time, location *time.Location) string {
	if location == nil {
		location = time.UTC
	}
	return t.In(location).Format("2006-01-02")
}
