package pushover

import (
	"fmt"
	"github.com/JamesClonk/go-todotxt"
	"github.com/gregdel/pushover"
	"time"
)

func SendNofification(token string, recipientKey string, task todotxt.Task) (err error) {
	app := pushover.New(token)
	recipient := pushover.NewRecipient(recipientKey)
	message := &pushover.Message{
		Message:   fmt.Sprintf("Task due on %s", task.DueDate.Format("2006-01-02")),
		Title:     task.Todo,
		Priority:  pushover.PriorityNormal,
		Timestamp: time.Now().Unix(),
		Retry:     60 * time.Second,
		Expire:    time.Hour,
	}

	_, err = app.SendMessage(message, recipient)
	return err
}
