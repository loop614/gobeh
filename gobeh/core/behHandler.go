package core

import (
	"github.com/loop614/gobeh/persistence"
	"log"
	"net/http"
)

func Start() {
	log.Print("Prepare db...")
	if err := persistence.Prepare(); err != nil {
		log.Fatal(err)
	}

	log.Print("Listening 8000")
	log.Fatal(http.ListenAndServe("gobeh_backend:8000", goBehServeHTTP{}))
}

type goBehServeHTTP struct {
}

func (gb goBehServeHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db, err := persistence.Connect()
	defer db.Close()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	routes := GetRoutes()
	path, pathOk := routes[r.RequestURI]
	if !pathOk {
		log.Print("Path not found")
		return
	}
	method, methodOk := path[r.Method]
	if !methodOk {
		log.Print("Method not found")
		return
	}
	log.Print(r.Method + " -> " + r.RequestURI)
	for key, val := range r.Form {
		log.Print("with" + key + "values: ")
		for _, oneVal := range val {
			log.Print(oneVal)
		}
	}
	method.handler(w, r, db)
}
