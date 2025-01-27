package types

import (
  "time"
)

type Task struct {
  Title string
  Description string
  DateCreated time.Time 
  Status Status 
}

type Status struct {
  NotStarted string
  InProgress string
  Completed string
}
