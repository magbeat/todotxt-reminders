package time

import (
	"time"
)

func NeedsNotification(reminder time.Time) (result bool, err error) {
	now := time.Now()
	if reminder.Before(now) {
		result = true
	}
	return result, err
}
