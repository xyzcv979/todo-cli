package main 

import (
	"log"
  "github.com/xyzcv979/todo-cli/internal/api"
	"github.com/xyzcv979/todo-cli/internal"
	"net/http"
)

func main() {
	log.Printf("server running on %s", internal.ServerURL)
	
	http.HandleFunc("/login", api.LoginHandler)
	http.HandleFunc("/createuser", api.CreateUserHandler)
	http.HandleFunc("/createtask", api.CreateTaskHandler)
	http.HandleFunc("/updatetask", api.UpdateTaskHandler)
	http.HandleFunc("/deletetask", api.DeleteTaskHandler)
	http.HandleFunc("/describetask", api.DescribeTaskHandler)
	http.HandleFunc("/listtask", api.ListTasksHandler)

	log.Print("Listening...")
	log.Fatalln(http.ListenAndServe(internal.ServerPort, nil))
}
