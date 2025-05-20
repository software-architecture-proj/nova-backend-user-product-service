package main

import (
	"fmt"
	"log"

    "google.golang.org/grpc"

	"github.com/software-architecture-proj/nova-backend-user-product-service/config"
	"github.com/software-architecture-proj/nova-backend-user-product-service/internal/repos"
	"github.com/software-architecture-proj/nova-backend-user-product-service/internal/server"	
    pb "github.com/software-architecture-proj/nova-backend-common-protos/gen/go/user_product_service"
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
    
    // Initialize repositories
    userRepo := repos.NewUserRepository(sqlDB)
	favoriteRepo := repos.NewFavoriteRepository(sqlDB)
	pocketRepo := repos.NewPocketRepository(sqlDB)

	// Create gRPC server instance
	service := &services.UserProductService{
		UserRepo:     userRepo,
		FavoriteRepo: favoriteRepo,
		PocketRepo:   pocketRepo,
	}
    
    // Listen TCP
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Iniciate gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterUserProductServiceServer(grpcServer, service)

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
