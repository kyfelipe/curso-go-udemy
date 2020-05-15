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

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")
	stmt.Exec(2000, "Santana")
	stmt.Exec(2001, "Carla")

	_, error := stmt.Exec(1, "Adriana") // chave duplicada
	if error != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
