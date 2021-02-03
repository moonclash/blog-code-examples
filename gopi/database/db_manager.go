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

// func getDB() {
// 	vardb, err := sql.Open("sqlite3", "./data.db")
// 	checkErr(err)
// 	return db, err
// }

func InsertDefinition() {

}

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