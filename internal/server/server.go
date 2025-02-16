package main 

import (
	"log"
  "github.com/xyzcv979/todo-cli/internal/api"
	"net/http"
)

func main() {
	log.Println("server running on http:localhost:8081")
	
	http.HandleFunc("/login", api.LoginHandler)
	http.HandleFunc("/createuser", api.CreateUserHandler)
	http.HandleFunc("/createtask", api.CreateTaskHandler)
	http.HandleFunc("/updatetask", api.UpdateTaskHandler)
	http.HandleFunc("/deletetask", api.DeleteTaskHandler)
	http.HandleFunc("/describetask", api.DescribeTaskHandler)
	http.HandleFunc("/listtask", api.ListTasksHandler)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
