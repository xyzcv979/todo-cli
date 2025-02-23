# todo-cli

Todo-cli is a CLI application built in Golang to create and track TODO tasks for improved productivity.
This project runs a server with APIs where clients send requests to. Clients first have to create an account with the server, login and authenticate, and finally can call CRUD operations on resources balonging to their user account.

Users will interact with the APIs through CLI Commands

### CLI 
Commands:
- Create 
- Delete
- Update
- List
- Login

### API endpoints
- "/login" 
- "/createuser" 
- "/createtask" 
- "/updatetask" 
- "/deletetask" 
- "/describetask"
- "/listtask"

### Models
Tasks:
- Name
- Date created
- Description
- Status

Status:
- Not started
- In progress
- Completed

## Components
Storage: Tasks and Users will be stored on a local sqliteDB file.
Authentication: JWT token signed when a client logins with username and password.
Authorization: Server will validate each incoming request for a valid token in the request headers.



