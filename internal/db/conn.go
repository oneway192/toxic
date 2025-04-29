package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "fmt"
)

var DB *sql.DB

// Init инициализация подключения к базе данных
func Init() error {
    var err error
    DB, err = sql.Open("postgres", "user=myuser password=toxic123 dbname=myappdb sslmode=disable")

    if err != nil {
        return err
    }

    err = DB.Ping()
    if err != nil {
        return err
    }
    log.Println("Connected to DB")
    return nil
}

// GetDB возвращает подключение к базе данных
func GetDB() (*sql.DB, error) {
    if DB == nil {
        return nil, fmt.Errorf("Database not initialized")
    }
    return DB, nil
}

