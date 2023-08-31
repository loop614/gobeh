package core

import (
	"database/sql"
	"github.com/loop614/gobeh/home"
	"net/http"
)

type Route struct {
	handler     func(w http.ResponseWriter, r *http.Request, db *sql.DB)
	handlerName string
}

func GetRoutes() map[string]map[string]Route {
	routes := map[string]map[string]Route{
		"/":    {"GET": {handler: home.Index, handlerName: "home.Index"}},
		"/add": {"POST": {handler: home.Add, handlerName: "home.Add"}},
	}

	return routes
}
