package time

import (
	"log"
	"time"
)

func NeedsNotification(reminderString string) (result bool, err error) {
	now := time.Now()
	zone, _ := now.Zone()
	reminder, err := time.Parse("2006-01-02T15:05 MST", reminderString + " " + zone)
	if err != nil {
		log.Printf("Could not parse date %v\n", reminderString)
	}
	if reminder.Before(now) {
		result = true
	}
	return result, err
}
