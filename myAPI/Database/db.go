package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres",
		"postgres://peoqxscq:o8KzOLhBc8U2tOjVkXN3g2Aj4iVSARXq@satao.db.elephantsql.com:5432/peoqxscq")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	log.Println("OK")
}
