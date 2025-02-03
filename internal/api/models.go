package api

// Define request and response models
import (
  "time"
)


type User struct {
	userId   int
	userName int
}

type Task struct {
	taskId      int
	userId      int
  Title       string
  Description string
  DateCreated time.Time 
  Status      Status 
}

type Status struct {
  NotStarted string
  InProgress string
  Completed  string
}
