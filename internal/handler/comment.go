package handler

import (
	"encoding/json"
	"log"
	"myapp/internal/db"
	"myapp/internal/model"
	"net/http"
	"strconv"
	"strings"
)

type CommentRequest struct {
	Text   string `json:"text_comment"`
	UserID int    `json:"user_id"`
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	var req CommentRequest
	json.NewDecoder(r.Body).Decode(&req)

	result, err := model.PredictText(req.Text)
	if err != nil {
		http.Error(w, "Model error", http.StatusInternalServerError)
		return
	}

	corrected := false
	if result == "CORRECTED" {
		corrected = true
	}

	_, err = db.DB.Exec("INSERT INTO comments (text_comment, user_id, corrected_ai) VALUES ($1, $2, $3)",
		req.Text, req.UserID, corrected)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, text_comment, user_id, corrected_ai FROM comments")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var comments []map[string]interface{}
	for rows.Next() {
		var id, userID int
		var text string
		var corrected bool
		rows.Scan(&id, &text, &userID, &corrected)

		comments = append(comments, map[string]interface{}{
			"id":        id,
			"text":      text,
			"user_id":   userID,
			"corrected": corrected,
		})
	}

	json.NewEncoder(w).Encode(comments)
}

func GetCommentsByUser(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	userID, _ := strconv.Atoi(parts[2])
	rows, err := db.DB.Query("SELECT id, text_comment, corrected_ai FROM comments WHERE user_id=$1", userID)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var comments []map[string]interface{}
	for rows.Next() {
		var id int
		var text string
		var corrected bool
		rows.Scan(&id, &text, &corrected)

		comments = append(comments, map[string]interface{}{
			"id":        id,
			"text":      text,
			"corrected": corrected,
		})
	}

	json.NewEncoder(w).Encode(comments)
}
