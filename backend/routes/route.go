package routes

import (
	"github.com/gorilla/mux"
	"backend/handlers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/emp", handlers.GetEmployeeHandler).Methods("GET")
	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
}
