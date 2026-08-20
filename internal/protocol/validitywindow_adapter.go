package protocol

import "time"

func ValidityWindowSameDay(left, right time.Time, location *time.Location) bool {
	return ValidityWindowDay(left, location) == ValidityWindowDay(right, location)
}
func ValidityWindowExpiry(t time.Time, location *time.Location) time.Time {
	if location == nil {
		location = time.UTC
	}
	y, m, d := t.In(location).Date()
	return time.Date(y, m, d, 23, 59, 59, 0, location)
}
