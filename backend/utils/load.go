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
		empData["employeeid"] = empID
		empJSON, _ := json.Marshal(empData)
		err := redis.SetValue(empID, string(empJSON))
		if err != nil {
			log.Printf("Failed to store employee %s: %v", empID, err)
		}
	}

	fmt.Println("Employees stored in Redis successfully!")
}

func SortCombinedToRedis() {
	var combined []Employee
	for i := 1; i <= 500; i++ {
		empID := fmt.Sprintf("EMP%04d", i)
		empData, err := redis.GetValue(empID)
		if err == redis.RedisNil {
			continue
		} else if err != nil {
			log.Printf("Failed to retrieve employee %s: %v", empID, err)
			continue
		}

		var emp Employee
		if err := json.Unmarshal([]byte(empData), &emp); err != nil {
			log.Printf("Failed to unmarshal employee %s: %v", empID, err)
			continue
		}

		combined = append(combined, emp)
	}
	combinedJSON, err := json.Marshal(combined)
	if err != nil {
		log.Printf("Failed to marshal combined employees: %v", err)
		return
	}
	err = redis.SetValue("employees", string(combinedJSON))
	if err != nil {
		log.Printf("Failed to store combined employees: %v", err)
	}
	fmt.Println("Combined employees stored in Redis successfully!")
}
