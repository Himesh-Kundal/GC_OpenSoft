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
	"backend/database"
)

func main() {
	config.LoadEnv()

	redis.InitRedis()
	utils.LoadEmployeesToRedis("./employee.json")
	go utils.SortCombinedToRedis()

	database.InitDB()

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	port := config.GetEnv("PORT", "8080")
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
