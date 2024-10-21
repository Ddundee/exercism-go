package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	time, _ := time.Parse("1/2/2006 15:04:05" ,date)
	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	return parsedTime.Before(time.Now())

}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	hour := parsedTime.Hour()
	return hour >= 12 && hour < 18


}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s, at %02d:%02d.",
		parsedTime.Format("Monday, January 2, 2006"), parsedTime.Hour(), parsedTime.Minute())

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	anniversary := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary

}
