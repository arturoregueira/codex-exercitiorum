package booking //bazinga

import (
	"fmt"
	"strconv"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	output, err := time.Parse(layout, date)
	if err == nil {
		return output
	} else {
		panic(fmt.Sprintf("Bad date %s was inputted", date))
	}

} // => Schedule("7/25/2019 13:45:00")

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	dateTime, err := time.Parse(layout, date)

	if err == nil {
		return dateTime.Before(time.Now())
	} else {
		panic(fmt.Sprintf("Bad date %s was inputted", date))
	}

} //HasPassed("July 25, 2019 13:45:00")
// => true

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	dateTime, err := time.Parse(layout, date)
	hourInt := dateTime.Hour()

	if err == nil {
		return hourInt >= 12 && hourInt < 18
	} else {
		panic(fmt.Sprintf("Bad date %s was inputted", date))
	}
} //IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
// => true

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	dateTime, err := time.Parse(layout, date)

	if err == nil {
		formatedDate := dateTime.Format("Monday, January 2, 2006, at 15:04.")
		return "You have an appointment on " + formatedDate
	} else {
		panic(fmt.Sprintf("Bad date %s was inputted", date))
	}
} //Description("7/25/2019 13:45:00")
// => "You have an appointment on Thursday, July 25, 2019, at 13:45."

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := strconv.Itoa(time.Now().Year())
	date := fmt.Sprintf("%s-09-15 00:00:00 +0000 UTC", year)
	layout := "2006-01-02 15:04:05 -0700 MST"
	dateTime, err := time.Parse(layout, date)

	if err == nil {

		return dateTime
	} else {
		panic(fmt.Sprintf("Bad date %s was inputted", date))
	}

} // => 2020-09-15 00:00:00 +0000 UTC
