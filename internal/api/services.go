package api

// Implement business logic.
// Interact with data models

import (
	"github.com/google/uuid"
)

func CreateUser() error{
	id := uuid.New()
  	
	return nil
}

func CreateTask(task Task) error {
  // store in database
  return nil
}

