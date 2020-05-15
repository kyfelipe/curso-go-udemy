package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:pass123@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, `create table if not exists usuarios (
		id auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
