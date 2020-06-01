# Send Pushover Notifications for Todo.txt Tasks

## Installation

```
go get github.com/magbeat/todotxt-reminders/cmd/todotxt-reminders
```

Set up an cron task, eg. every 5 minutes

```
*/5 * * * * todotxt-reminders <path to todo.txt file> <pushover app token> <pushover device token>
```

## Add reminder to a task

```
2020-05-31 myTask reminder:2020-06-01T11:00 due:2020-06-01
```

