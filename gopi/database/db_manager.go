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


type DBManager struct {
	name string
	db *sql.DB
}

func new(name string) DBManager {
	_db := DBManager {name, nil}
	return _db
}

func (_db DBManager) initialize () {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	_db.db = db
}

func (_db DBManager) InsertDefinition(key string, definition string) {
	insertQuery := `INSERT INTO definitions(short_name, long_name) VALUES (?, ?)`
	statement, err := _db.db.Prepare(insertQuery)
	checkErr(err)
	_, err = statement.Exec(key, definition)
	checkErr(err)
}




// db, err := sql.Open("sqlite3", "./data.db")
// checkErr(err)

// func InsertDefinition(db *sql.DB, key string, definition string) {
// 	insertQuery := `INSERT INTO definitions(short_name, long_name) VALUES (?, ?)`
// 	statement, err := db.Prepare(insertQuery)
// 	checkErr(err)
// 	_, err = statement.Exec(key, definition)
// 	checkErr(err) 
// }

// func RetrieveDefinition() {
// 	db := getDB()
// }

func DeleteDefinition() {

}

func CreateDB() {
	os.Create("./data.db")
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `definitions` (`definition_id` INTEGER PRIMARY KEY AUTOINCREMENT, `short_name` VARCHAR(255), `long_name` VARCHAR(255))")
	checkErr(err)
	db.Close()
}