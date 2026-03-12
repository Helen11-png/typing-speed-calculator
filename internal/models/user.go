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

// Функция для получения профиля пользователя
func GetUserProfile(userID int) (map[string]interface{}, error) {
	var username string
	var createdAt string
	var testsCount int
	var avgSpeed float64
	var bestSpeed int

	query := `
        SELECT 
            u.username,
            u.created_at,
            COUNT(r.id) as tests_count,
            COALESCE(AVG(r.speed), 0) as avg_speed,
            COALESCE(MAX(r.speed), 0) as best_speed
        FROM users u
        LEFT JOIN results r ON u.id = r.user_id
        WHERE u.id = ?
        GROUP BY u.id, u.username, u.created_at
    `

	err := DB.QueryRow(query, userID).Scan(
		&username, &createdAt, &testsCount, &avgSpeed, &bestSpeed,
	)

	if err != nil {
		return nil, err
	}

	profile := map[string]interface{}{
		"username":    username,
		"created_at":  createdAt,
		"tests_count": testsCount,
		"avg_speed":   avgSpeed,
		"best_speed":  bestSpeed,
	}

	return profile, nil
}
