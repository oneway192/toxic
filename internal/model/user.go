package model

// User структура для пользователя
type User struct {
    ID       int    `json:"id"`
    Login    string `json:"login"`
    Password string `json:"password"`
}

