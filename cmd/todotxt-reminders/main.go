package main

import (
	"github.com/JamesClonk/go-todotxt"
	"github.com/magbeat/todotxt-reminders/internal/pushover"
	time2 "github.com/magbeat/todotxt-reminders/internal/time"
	"log"
	"os"
	"time"
)

func main() {
	todotxt.IgnoreComments = false

	if len(os.Args) != 4 {
		log.Fatal("Needs exactly two arguments: the todo.txt path and the time span in minutes")
	}
	todoListFilename := os.Args[1]
	pushoverToken := os.Args[2]
	pushoverDevice := os.Args[3]

	taskList, err := todotxt.LoadFromFilename(todoListFilename)
	if err != nil {
		log.Fatal(err)
	}

	for _, task := range taskList {
		reminder, reminderSet := task.AdditionalTags["reminder"]
		if reminderSet {
			t, err := time.Parse("2006-01-02T15:05", reminder)
			if err != nil {
				log.Println("Could not parse date %v", reminder)
			}

			needsNotification, err := time2.NeedsNotification(t)
			if err != nil {
				log.Println(err)
			}

			if needsNotification {
				err = pushover.SendNofification(pushoverToken, pushoverDevice, task)
				if err == nil {
					delete(task.AdditionalTags, "reminder")
					err = taskList.WriteToFilename(todoListFilename)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}
