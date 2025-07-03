package storage

import (
	"database/sql"
	"log"
	"os"
    _ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS leads (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fullname TEXT NOT NULL,
		position TEXT NOT NULL,
		company TEXT NOT NULL,
		email TEXT NOT NULL,
		phone TEXT NOT NULL,
		course_description TEXT,
		time TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}
}
