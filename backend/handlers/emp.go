package handlers

import (
	"context"
	"net/http"
	backendRedis "backend/redis"
	"github.com/go-redis/redis/v8"
)

// Get employee details from Redis
func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	empID := r.URL.Query().Get("id")
	if empID == "" {
		http.Error(w, "Employee ID is required", http.StatusBadRequest)
		return
	}

	empData, err := backendRedis.RedisClient.Get(context.Background(), empID).Result()
	if err == redis.Nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error retrieving employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(empData))
}

// Health check handler
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := backendRedis.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		http.Error(w, "Redis is unavailable", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
