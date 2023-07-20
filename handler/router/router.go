package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	//mux.HandleFunc("http://localhost:8080/healthz", handler.NewHealthzHandler())
	mux.Handle("/healthz", handler.NewHealthzHandler())

	//todo
	mux.Handle("http://localhost:8080/todos", handler.NewTODOHandler(service.NewTODOService(todoDB)))

	return mux
}
