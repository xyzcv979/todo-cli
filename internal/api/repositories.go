package api

// Setup database interaction

/*
Table schemas:

CREATE TABLE if not exists user(
userId INTEGER PRIMARY KEY,
userName TEXT);

CREATE TABLE if not exists task (
taskId INTEGER PRIMARY KEY,
userId INTEGER,
Title TEXT NOT NULL,
Description TEXT,
DateCreated TEXT NOT NULL,
Status TEXT NOT NULL,
FOREIGN KEY (userId) REFERENCES User(userId) ON DELETE CASCADE);

*/
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"log"
)

const (
	dbName = "../../todoDB.db" 
  taskTable = `
		CREATE TABLE if not exists task (
		taskId INTEGER PRIMARY KEY,
		userId INTEGER,
		Title TEXT NOT NULL,
		Description TEXT,
		DateCreated TEXT NOT NULL,
		Status TEXT NOT NULL,
		FOREIGN KEY (userId) REFERENCES User(userId) ON DELETE CASCADE);`  
	userTable = `
		CREATE TABLE if not exists user(
		userId INTEGER PRIMARY KEY,
		userName TEXT);`
)

func connectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)	
	}
	defer db.Close()
	
	_, err = db.Exec(userTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(taskTable)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to sqlite database")
	return db
}

func InsertUser(db *sql.DB, user User) error {
	_, err := db.Exec("INSERT INTO user (userName) VALUES (?)", user.userName)
	if err != nil {
		return err
	}
	return nil	
}

func InsertTask(db *sql.DB, task Task) error {
	_, err := db.Exec("INSERT INTO task (userId, Title, Desription, DateCreated, Status) (?, ?, ?, ?, ?)", task.userId, task.Title, task.Description, task.DateCreated, task.Status)
	if err != nil {
		return err
	}
	return nil	
}
