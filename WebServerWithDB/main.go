package main

import (
	"context"
	"database-example/db"
	"database-example/handler"
	"database-example/model"
	follower_service "database-example/proto/follower"
	"database-example/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

type Server struct {
	follower_service.UnimplementedFollowerServiceServer
	FollowerService *service.FollowerService
}

func NewServer(driver neo4j.Driver) *Server {
	follower_service := service.NewFollowerService(driver)
	return &Server{
		FollowerService: follower_service,
	}
}
func startServer() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	router := mux.NewRouter().StrictSlash(true)

	followerHandler := handler.NewFollowerHandler(database)
	followerHandler.RegisterRoutes(router)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	log.Println("Server is running on port", db.Port)
	log.Fatal(http.ListenAndServe(":8082", router))
}

func ConvertToIntSlice(int32Slice []int32) []int {
	intSlice := make([]int, len(int32Slice))
	for i, v := range int32Slice {
		intSlice[i] = int(v)
	}
	return intSlice
}

func ConvertToInt32Slice(intSlice []int) []int32 {
	int32Slice := make([]int32, len(intSlice))
	for i, v := range intSlice {
		int32Slice[i] = int32(v)
	}
	return int32Slice
}

func (s *Server) UpsertFollower(ctx context.Context, req *follower_service.UpsertFollowerRequest) (*follower_service.UpsertFollowerResponse, error) {
	follower := &model.Follower{
		Id:         req.Follower.GetId(),
		Followers:  ConvertToIntSlice(req.Follower.GetFollowers()),
		Followable: ConvertToIntSlice(req.Follower.GetFollowable()),
		Followed:   ConvertToIntSlice(req.Follower.GetFollowed()),
	}

	err := s.FollowerService.CreateFollower(*follower)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert user: %v", err)
	}

	return &follower_service.UpsertFollowerResponse{
		Follower: &follower_service.Follower{
			Id:         follower.Id,
			Followers:  ConvertToInt32Slice(follower.Followers),
			Followable: ConvertToInt32Slice(follower.Followable),
			Followed:   ConvertToInt32Slice(follower.Followed),
		},
	}, nil
}

func (s *Server) GetFollower(ctx context.Context, req *follower_service.GetFollowerRequest) (*follower_service.GetFollowerResponse, error) {
	follower, err := s.FollowerService.GetById(int(req.GetId()))
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return &follower_service.GetFollowerResponse{
		Follower: &follower_service.Follower{
			Id:         follower.Id,
			Followers:  ConvertToInt32Slice(follower.Followers),
			Followable: ConvertToInt32Slice(follower.Followable),
			Followed:   ConvertToInt32Slice(follower.Followed),
		},
	}, nil
}

func main() {

	//startServer()

	driver, err := db.InitDB()
	if err != nil {
		log.Fatal("FAILED TO CONNECT TO NEO4J")
	}
	defer driver.Close()

	// Inicijalizacija servera
	server := NewServer(driver)

	// Inicijalizacija gRPC servera
	grpcServer := grpc.NewServer()
	follower_service.RegisterFollowerServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	// Slušanje na TCP portu
	lis, err := net.Listen("tcp", ":50053")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server listening on port 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func (s *Server) DeleteFollower(ctx context.Context, req *follower_service.DeleteFollowerRequest) (*follower_service.DeleteFollowerResponse, error) {
	err := s.FollowerService.DeleteFollower(int(req.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to delete encounter: %v", err)
	}
	return &follower_service.DeleteFollowerResponse{}, nil
}