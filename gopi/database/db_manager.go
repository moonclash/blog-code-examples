package database

import (
	"os"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CreateDB() {
	os.Create("./data.db")
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `definitions` (`definition_id` INTEGER PRIMARY KEY AUTOINCREMENT, `short_name` VARCHAR(255), `long_name` VARCHAR(255))")
	checkErr(err)
}

type DBManager struct {
	db *sql.DB
}

func New(db *sql.DB) DBManager {
	_db := DBManager {db}
	return _db
}

func (_db DBManager) Initialize() {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	_db.db = db
}

func (_db DBManager) InsertDefinition(key string, definition string) {
	insertQuery := "INSERT INTO definitions (short_name, long_name) VALUES (?, ?)"
	statement, err := _db.db.Prepare(insertQuery)
	checkErr(err)
	statement.Exec(key, definition)
	checkErr(err)
}

func (_db DBManager) RetrieveDefinition(key string) {
	string_query := fmt.Sprintf("select short_name, long_name from definitions where short_name = %s", key)
	rows, err := _db.db.Query(string_query)
	fmt.Println(rows)
	checkErr(err)
}