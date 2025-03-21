package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"backend/redis"
)

type Employee map[string]interface{}

func LoadEmployeesToRedis(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read employee.json: %v", err)
	}

	var employees map[string]Employee
	if err := json.Unmarshal(file, &employees); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for empID, empData := range employees {
		empJSON, _ := json.Marshal(empData)
		err := redis.SetValue(empID, string(empJSON))
		if err != nil {
			log.Printf("Failed to store employee %s: %v", empID, err)
		}
	}

	fmt.Println("Employees stored in Redis successfully!")
}
