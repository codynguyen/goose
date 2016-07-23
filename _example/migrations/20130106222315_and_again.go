package main

import (
	"database/sql"
	"fmt"
)

func Up_20130106222315(txn *sql.Tx, db *sql.DB) {
	fmt.Println("Hello from migration 20130106222315 Up!")
}

func Down_20130106222315(txn *sql.Tx, db *sql.DB) {
	fmt.Println("Hello from migration 20130106222315 Down!")
}
