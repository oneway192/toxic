package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

// Init инициализирует подключение к базе данных
func Init() error {
    var err error
    DB, err = sql.Open("postgres", "user=myuser password=toxic123 dbname=myappdb sslmode=disable")
    if err != nil {
        return fmt.Errorf("failed to open DB: %w", err)
    }

    if err := DB.Ping(); err != nil {
        return fmt.Errorf("failed to connect to DB: %w", err)
    }

    log.Println("Connected to DB")
    return nil
}

// GetDB возвращает текущее подключение к базе данных
func GetDB() (*sql.DB, error) {
    if DB == nil {
        return nil, fmt.Errorf("database not initialized")
    }
    return DB, nil
}

