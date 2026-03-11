package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./data/app.db")
	if err != nil {
		return err
	}

	// Создаем таблицы
	createTables := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT UNIQUE,
        password TEXT,
        created_at DATETIME
    );
    
    CREATE TABLE IF NOT EXISTS results (
        id INTEGER PRIMARY KEY,
        user_id INTEGER,
        speed INTEGER,
        accuracy INTEGER,
        text_preview TEXT,
        created_at DATETIME,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`

	_, err = DB.Exec(createTables)
	if err != nil {
		return err
	}

	log.Println("База данных инициализирована")
	return nil
}
