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

	if err := db.Migrate(); err != nil {
		log.Fatal("Migration error:", err)
	}

	http.HandleFunc("/register", handler.RegisterUser)
	http.HandleFunc("/login", handler.LoginUser)
	http.HandleFunc("/comment", handler.AddComment)
	http.HandleFunc("/comments", handler.GetAllComments)
	http.HandleFunc("/comments/", handler.GetCommentsByUser)

	http.HandleFunc("/users", handler.GetAllUsers)

	log.Println("Server started at :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
