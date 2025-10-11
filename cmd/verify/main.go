package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:1234name@localhost:5432/expense_tracker?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("База данных недоступна:", err)
	}

	fmt.Println("✅ Успешное подключение к PostgreSQL!")

	rows, err := db.Query(`
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema='public' 
		ORDER BY table_name;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("📋 Таблицы в базе данных:")
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Printf(" - %s\n", name)
	}
}
