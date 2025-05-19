package server

import (
    "context"
	"log"
	"net"
	"os"
    "strconv"

pb "github.com/software-architecture-proj/nova-backend-common-protos/gen/go/user_product_service" 
	"google.golang.org/grpc"
    "github.com/google/uuid"
	"google.golang.org/grpc/reflection"
    "github.com/software-architecture-proj/nova-backend-user-product-service/internal/models"
    "github.com/software-architecture-proj/nova-backend-user-product-service/config"
)

// userProductServer is the implementation of the UserServiceServer interface.
type userProductServer struct {
	pb.UnimplementedUserServiceServer // Embed this to handle unimplemented methods.
}

// CreateUser implements the CreateUser RPC method.
func (s *userProductServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
    createdUser := &models.User{
        ID:         uuid.New(),
        Email:      req.GetEmail(),
        Username:   req.GetUsername(),
        Phone:      strconv.ParseInt(req.GetPhone(), 10, 64),
        //@trigger: Don't know what to do here...
    }

	return newUser, nil
}

// GetUserById implements the GetUserById RPC method.
func (s *userProductServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	// Implement your user retrieval logic here.
	log.Printf("Received GetUserById request: %v", req)

	// Example (replace with database lookup):
	if req.UserId == "user-id-123" {
		user := &pb.User{
			Id:        "user-id-123",
			Email:     "test@example.com",
			Username:  "testuser",
			Phone:     "123-456-7890",
			FirstName: "Test",
			LastName:  "User",
			Birthdate: "1990-01-01",
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		}
		return user, nil
	}
	return nil, status.Errorf(codes.NotFound, "User not found") //  Use status.Errorf
}

// UpdateUserById implements the UpdateUserById RPC method.
func (s *userProductServer) UpdateUserById(ctx context.Context, req *pb.UpdateUserByIdRequest) (*pb.User, error) {
	// Implement your user update logic here
	log.Printf("Received UpdateUserById request: %v", req)
    updatedUser := &pb.User{
		Id:        req.Id,
		Email:     req.Email,
		Username:  req.Username,
		Phone:     req.Phone,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Birthdate: req.Birthdate,
		UpdatedAt: time.Now().String(),
	}
	return updatedUser, nil
}

// DeleteUserById implements the DeleteUserById RPC.
func (s *userProductServer) DeleteUserById(ctx context.Context, req *pb.DeleteUserByIdRequest) (*pb.DeleteResponse, error) {
	// Implement your delete logic here.
	log.Printf("Received DeleteUserById request: %v", req)
	return &pb.DeleteResponse{Success: true}, nil
}

// CreateFavorite implements the CreateFavorite RPC method.
func (s *userProductServer) CreateFavorite(ctx context.Context, req *pb.CreateFavoriteRequest) (*pb.Favorite, error) {
    createdFavorite := &models.Favorite{
        ID:             uuid.New(),
        UserID:         req.GetUserId(),
        //User: @trigger: I don't understand, the struct should have an User struct inside?
        FavoriteUserID: req.GetFavoriteUserId(),
        Alias:          req.GetAlias(),
    }

    repo := NewFavoriteRepository(DB)
    repo.CreateFavorite(createdFavorite)

	return newFavorite, nil
}

func (s *userProductServer) GetFavoritesByUserId(ctx context.Context, req *pb.GetFavoritesByUserIdRequest) (*pb.GetFavoritesByUserIdResponse, error) {
	// Implement
	log.Printf("Received GetFavoritesByUserId: %v", req)
	favorites := []*pb.Favorite{
		{
			Id:             "fav-id-1",
			UserId:         req.UserId,
			FavoriteUserId: "user-id-2",
			Alias:          "My Favorite User",
			CreatedAt:      time.Now().String(),
			UpdatedAt:      time.Now().String(),
		},
	}
	return &pb.GetFavoritesByUserIdResponse{Favorites: favorites}, nil
}

func (s *userProductServer) UpdateFavoriteById(ctx context.Context, req *pb.UpdateFavoriteByIdRequest) (*pb.Favorite, error) {
	// Implement
	log.Printf("Received UpdateFavoriteById: %v", req)
	updatedFavorite := &pb.Favorite{
		Id:             req.Id,
		Alias:          req.Alias,
		UpdatedAt:      time.Now().String(),
	}
	return updatedFavorite, nil
}

func (s *userProductServer) DeleteFavoriteById(ctx context.Context, req *pb.DeleteFavoriteByIdRequest) (*pb.DeleteResponse, error) {
	// Implement
	log.Printf("Received DeleteFavoriteById: %v", req)
	return &pb.DeleteResponse{Success: true}, nil
}

// CreatePocket implements the CreatePocket RPC.
func (s *userProductServer) CreatePocket(ctx context.Context, req *pb.CreatePocketRequest) (*pb.Pocket, error) {
	log.Printf("Received CreatePocket: %v", req)
	pocket := &pb.Pocket{
		Id:           "pocket-1", // generate
		UserId:       req.UserId,
		PocketUserId: req.PocketUserId,
		Alias:        req.Alias,
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}
	return pocket, nil
}

// GetPocketsByUserId implements the GetPocketsByUserId RPC.
func (s *userProductServer) GetPocketsByUserId(ctx context.Context, req *pb.GetPocketsByUserIdRequest) (*pb.GetPocketsByUserIdResponse, error) {
	log.Printf("Received GetPocketsByUserId: %v", req)
	pockets := []*pb.Pocket{
		{
			Id:           "pocket-1",
			UserId:       req.UserId,
			PocketUserId: "user-2",
			Alias:        "My Pocket",
			CreatedAt:    time.Now().String(),
			UpdatedAt:    time.Now().String(),
		},
	}
	return &pb.GetPocketsByUserIdResponse{Pockets: pockets}, nil
}

// UpdatePocketById implements the UpdatePocketById RPC.
func (s *userProductServer) UpdatePocketById(ctx context.Context, req *pb.UpdatePocketByIdRequest) (*pb.Pocket, error) {
	log.Printf("Received UpdatePocketById: %v", req)
	pocket := &pb.Pocket{
		Id:        req.Id,
		Alias:     req.Alias,
		UpdatedAt: time.Now().String(),
	}
	return pocket, nil
}

// DeletePocketById implements the DeletePocketById RPC.
func (s *userProductServer) DeletePocketById(ctx context.Context, req *pb.DeletePocketByIdRequest) (*pb.DeleteResponse, error) {
	log.Printf("Received DeletePocketById: %v", req)
	return &pb.DeleteResponse{Success: true}, nil
}
