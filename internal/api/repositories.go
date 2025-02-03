package api

// Setup database interaction

/*
Table schemas:

CREATE TABLE user(
userId INTEGER PRIMARY KEY,
userName TEXT);

CREATE TABLE task (
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
)

func connectToDB() {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)	
	}
	defer db.Close()

	log.Println("Connected to sqlite database")
}
func AddUsertoDB(user User) {

}
