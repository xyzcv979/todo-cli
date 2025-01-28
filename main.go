package main

import (
  "log"
  "fmt"
  "github.com/xyzcv979/todo-cli/types"
  "github.com/spf13/cobra"
)

func main() {
  log.SetFlags(log.LstdFlags | log.Lshortfile)
  log.Println("hello world") 

  firstTask := types.Task{}
  fmt.Println(firstTask)
}
