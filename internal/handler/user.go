package handler

import (
    "encoding/json"
    "log"
    "net/http"
    "myapp/internal/db"
    "myapp/internal/model"
)

// RegisterUser обрабатывает POST-запрос на регистрацию нового пользователя
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user model.User

    // Декодируем тело запроса в структуру User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        log.Printf("Error decoding input: %v", err)
        return
    }

    // Подключаемся к базе данных
    db, err := db.GetDB()
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        log.Printf("Database connection error: %v", err)
        return
    }

    // Проверяем, существует ли уже пользователь с таким логином
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE login = $1", user.Login).Scan(&count)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        log.Printf("Error checking login uniqueness: %v", err)
        return
    }

    if count > 0 {
        http.Error(w, "User with this login already exists", http.StatusConflict)
        return
    }

    // SQL запрос на добавление нового пользователя
    query := "INSERT INTO users (login, password) VALUES ($1, $2)"
    _, err = db.Exec(query, user.Login, user.Password)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        log.Printf("Error executing query: %v", err)
        return
    }

    // Ответ клиенту о успешной регистрации
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("User registered successfully")
}

// LoginUser обрабатывает POST-запрос на логин пользователя
func LoginUser(w http.ResponseWriter, r *http.Request) {
    var user model.User

    // Декодируем тело запроса
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        log.Printf("Error decoding input: %v", err)
        return
    }

    // Подключаемся к базе данных
    db, err := db.GetDB()
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        log.Printf("Database connection error: %v", err)
        return
    }

    // Проверяем, существует ли пользователь с таким логином и паролем
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE login = $1 AND password = $2", user.Login, user.Password).Scan(&count)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        log.Printf("Error checking user credentials: %v", err)
        return
    }

    if count == 0 {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Ответ клиенту о успешном входе
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Login successful")
}

// GetAllUsers обрабатывает GET-запрос на получение всех пользователей
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    db, err := db.GetDB()
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        log.Printf("Database connection error: %v", err)
        return
    }

    // Выполняем запрос для получения всех пользователей
    rows, err := db.Query("SELECT id, login FROM users")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        log.Printf("Error fetching users: %v", err)
        return
    }
    defer rows.Close()

    var users []map[string]interface{}
    for rows.Next() {
        var id int
        var login string
        if err := rows.Scan(&id, &login); err != nil {
            http.Error(w, "Error processing query result", http.StatusInternalServerError)
            log.Printf("Error scanning row: %v", err)
            return
        }

        users = append(users, map[string]interface{}{
            "id":    id,
            "login": login,
        })
    }

    // Кодируем список пользователей в JSON и отправляем клиенту
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}

