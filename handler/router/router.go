package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	//mux.HandleFunc("http://localhost:8080/healthz", handler.NewHealthzHandler())
	mux.Handle("/healthz", handler.NewHealthzHandler())
	//mux.HandleFunc("http://localhost:8080/healthz", handler.HealthzHandler)
	return mux
}
