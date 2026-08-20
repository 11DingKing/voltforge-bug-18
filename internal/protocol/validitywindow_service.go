package protocol

import "time"

func ValidityWindowDay(t time.Time, location *time.Location) string {
	return t.UTC().Format("2006-01-02")
}
