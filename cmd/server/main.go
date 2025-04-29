package main

import (
    "log"
    "net/http"

    "myapp/internal/db"
    "myapp/internal/handler"
)

func main() {
    err := db.Init()
    if err != nil {
        log.Fatal("DB init error:", err)
    }

    http.HandleFunc("/register", handler.RegisterUser)
    http.HandleFunc("/login", handler.LoginUser)
    http.HandleFunc("/comment", handler.AddComment)
    http.HandleFunc("/comments", handler.GetAllComments)
    http.HandleFunc("/comments/", handler.GetCommentsByUser)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

