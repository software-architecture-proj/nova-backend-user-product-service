package main

import (
	"fmt"
	"log"
	"nova-backend-user-product-service/config"
)

func main() {
	config.InitDB()

	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatalf("Database instance error: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Database connection established.")

}
