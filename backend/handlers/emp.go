package handlers

import (
	"net/http"
	"backend/redis"
)

func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	empID := r.URL.Query().Get("id")
	if empID == "" {
		http.Error(w, "Employee ID is required", http.StatusBadRequest)
		return
	}
	// Missing Employees
	// -EMP0099
	// -EMP0240
	// -EMP0374
	empData, err := redis.GetValue(empID)
	if err ==  redis.RedisNil{
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error retrieving employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(empData))
}

func GetAllEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	empData, err :=redis.GetValue("employees")
	if err == redis.RedisNil {
		http.Error(w, "Employees not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error retrieving employees", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(empData))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := redis.Ping()
	if err != nil {
		http.Error(w, "Redis is unavailable", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
