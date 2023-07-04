package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	//mux.Handle("http://localhost:8080/healthz", http.HandleFunc(http.Handler.ServeHTTP()))
	//mux.Handle("http://localhost:8080/healthz", http.Handler.HealthzHandler)
	mux.HandleFunc("http://localhost:8080/healthz", handler.HealthzHandler)
	//mux.HandleFunc("http://localhost:8080/healthz", handler.HealthzHandler)
	return mux
}
