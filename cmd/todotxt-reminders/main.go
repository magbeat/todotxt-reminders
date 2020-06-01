package main

import (
	"fmt"
	"github.com/JamesClonk/go-todotxt"
	"github.com/magbeat/todotxt-reminders/internal/pushover"
	time2 "github.com/magbeat/todotxt-reminders/internal/time"
	"log"
	"os"
)

func main() {
	todotxt.IgnoreComments = false

	if len(os.Args) != 4 {
		log.Fatal("Needs exactly 3 arguments: the todo.txt file path, the pushover app token and the pushover device token")
	}
	todoListFilename := os.Args[1]
	pushoverToken := os.Args[2]
	pushoverDevice := os.Args[3]

	taskList, err := todotxt.LoadFromFilename(todoListFilename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %v tasks\n", len(taskList))
	tasksWithReminders := 0
	notificationsSent := 0
	for _, task := range taskList {
		reminder, reminderSet := task.AdditionalTags["reminder"]
		if reminderSet {
			tasksWithReminders++
			needsNotification, err := time2.NeedsNotification(reminder)
			if err != nil {
				log.Println(err)
			}

			if needsNotification {
				err = pushover.SendNofification(pushoverToken, pushoverDevice, task)
				if err == nil {
					notificationsSent++
					delete(task.AdditionalTags, "reminder")
					err = taskList.WriteToFilename(todoListFilename)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}

	fmt.Printf("Found %v tasks with reminders\n", tasksWithReminders)
	fmt.Printf("Sent %v notifications\n", notificationsSent)
}
