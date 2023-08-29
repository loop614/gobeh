package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/loop614/gobeh/core"
	"github.com/loop614/gobeh/persistence"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Prepare db...")

	if err := persistence.Prepare(); err != nil {
		log.Fatal(err)
	}

	log.Print("Listening 8000")
	r := mux.NewRouter()
	r.HandleFunc("/", core.BlogHandler)
	log.Fatal(http.ListenAndServe("gobeh_backend:8000", handlers.LoggingHandler(os.Stdout, r)))
}
