package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// func main() {
// 	url := os.Getenv("DATABASE_URL")
// 	// url := "postgres://peoqxscq:o8KzOLhBc8U2tOjVkXN3g2Aj4iVSARXq@satao.db.elephantsql.com:5432/peoqxscq"
// 	db, err := sql.Open("postgres", url)

// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	defer db.Close()

// 	createTb := `
// 	CREATE TABLE IF NOT EXISTS todos (
// 		id SERIAL PRIMARY KEY,
// 		title TEXT,
// 		status TEXT
// 	);
// 	`

// 	// selectTb := `
// 	// SELECT *
// 	// from todos
// 	// `

// 	_, err = db.Exec(createTb)

// 	if err != nil {
// 		log.Fatal("can't create table", err)
// 	}

// 	log.Println("create table success")
// }

func createTable() {
	url := os.Getenv("DATABASE_URL")
	// url := "postgres://peoqxscq:o8KzOLhBc8U2tOjVkXN3g2Aj4iVSARXq@satao.db.elephantsql.com:5432/peoqxscq"
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	log.Println("create table success")
}

func insertTodo() {
	url := "postgres://peoqxscq:o8KzOLhBc8U2tOjVkXN3g2Aj4iVSARXq@satao.db.elephantsql.com:5432/peoqxscq"
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()
	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", "buy bmw", "active")
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo success id : ", id)
}

func queryOneRow() {
	// url := "postgres://peoqxscq:o8KzOLhBc8U2tOjVkXN3g2Aj4iVSARXq@satao.db.elephantsql.com:5432/peoqxscq"
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var id int
	var title, status string

	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	fmt.Println("one row", id, title, status)
}

func main() {
	// insertTodo()

	// os.Setenv("DATABASE_URL", "postgres://peoqxscq:o8KzOLhBc8U2tOjVkXN3g2Aj4iVSARXq@satao.db.elephantsql.com:5432/peoqxscq")
	fmt.Println(os.Getenv("DATABASE_URL"))

	// queryOneRow()
}
