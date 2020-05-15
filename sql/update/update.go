package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:pass123@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sheristone Uasleska", 2)
	stmt.Close()

	// Delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
	stmt2.Close()
}
