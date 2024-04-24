package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowerHandler struct {
	service *service.FollowerService
}

func NewUserHandler(driver neo4j.Driver) *FollowerHandler {
	followerService := service.NewFollowerService(driver)
	return &FollowerHandler{
		service: followerService,
	}
}

func (h *FollowerHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/follow", h.CreateUserHandler).Methods("POST")

}

func (uh *FollowerHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.Follower
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = uh.service.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
