package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func openConnection() (db *sql.DB, err error) {
	// url := "postgres://peoqxscq:o8KzOLhBc8U2tOjVkXN3g2Aj4iVSARXq@satao.db.elephantsql.com:5432/peoqxscq"
	url := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", url)
	return
}
func createTable() {
	db, err := openConnection()
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

	// db.Exec means to execute the given SQL
	// w/o any result|state from the result
	// following the execution
	// however the PLACEHOLDER is used to track
	// number of rows AFFECTED from the execution
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	log.Println("create table success")
}

func insertTodo() {
	db, err := openConnection()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id;", "talk with my ex", "never")
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo success id : ", id)
}

func queryOneRow() {
	db, err := openConnection()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1;")
	if err != nil {
		log.Fatal("can't prepare query one row statment", err)
	}

	rowId := 1

	// rowId corresponds to $variable in db.Prepare
	// can use multiple variables by adding more $variable
	row := stmt.QueryRow(rowId)
	var id int
	var title, status string

	// store results from stmt.QueryRow to each variable
	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	fmt.Println("one row", id, title, status)
}

func updateTodo() {
	db, err := openConnection()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("update todos set status=$1 where id=$2;")
	if err != nil {
		log.Fatal("can't prepare update statement", err)
	}

	if _, err := stmt.Exec("complete", 3); err != nil {
		log.Fatal("can't update", err)
	}
	fmt.Println("update success")
}

func queryAllTodo() {
	db, err := openConnection()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statement", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}

	var id int
	var title, status string
	for rows.Next() {
		if err := rows.Scan(&id, &title, &status); err != nil {
			log.Fatal("can't scane row into id,title,status", err)
			continue
		}
		fmt.Println(id, title, status)
	}
	fmt.Println("query all todos success")
}

func main() {
	// queryOneRow()
	// insertTodo()
	// updateTodo()
	queryAllTodo()
}
