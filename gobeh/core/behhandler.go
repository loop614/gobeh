package core

import (
	"encoding/json"
	"github.com/loop614/gobeh/persistence"
	"net/http"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	db, err := persistence.Connect()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer db.Close()

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
