package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"backend/config"
	"backend/redis"
	"backend/routes"
	"backend/utils"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize Redis client
	redis.InitRedis()

	// Load employees into Redis
	utils.LoadEmployeesToRedis("employee.json")

	// Set up router
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Start server
	port := config.GetEnv("PORT", "8080")
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
