package main

import (
	"fmt"
	"log"

	"github.com/software-architecture-proj/nova-backend-user-product-service/config"
)

func main() {
	log.Println("🔧 Starting user product service...")

	if err := config.InitDB(); err != nil {
		log.Fatalf("DB initialization failed: %v", err)
	}
	log.Println("DB initialized successfully")

	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatalf("Getting DB instance failed: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Ping to DB failed: %v", err)
	}

	fmt.Println("Database connection established.")
}
