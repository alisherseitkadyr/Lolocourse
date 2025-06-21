package storage

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS leads (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		phone TEXT NOT NULL,
		email TEXT,
		time TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}
}
