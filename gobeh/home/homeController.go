package home

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT title FROM blog")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var titles []string

	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		titles = append(titles, title)
	}
	json.NewEncoder(w).Encode(titles)
}

type Blogpost struct {
	Title string `json:"title"`
}

func Add(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var blogpost Blogpost
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&blogpost)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if _, err := db.Exec("INSERT INTO blog (title) VALUES ($1);", blogpost.Title); err != nil {
		w.WriteHeader(500)
		return
	}
	rows, err := db.Query("SELECT title FROM blog")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var titles []string

	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		titles = append(titles, title)
	}
	json.NewEncoder(w).Encode(titles)
}
