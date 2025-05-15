package main

/*

import (
    "log"
    "net"

    "nova-backend-user-product-service/config"
    "nova-backend-user-product-service/internal/user"
    "nova-backend-user-product-service/internal/user/pb"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

func main() {
    config.InitDB()

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    userRepo := user.NewRepository(config.DB)
    userService := user.NewService(userRepo)

    pb.RegisterUserServiceServer(grpcServer, userService)

    // Optional reflection for CLI tools
    reflection.Register(grpcServer)

    log.Println("gRPC server listening on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
*/
