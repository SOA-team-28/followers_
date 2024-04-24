package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"

	"net/http"
)

type FollowerHandler struct {
	followerService *service.FollowerService
}

func NewUserHandler(followerService *service.FollowerService) *FollowerHandler {
	return &FollowerHandler{
		followerService: followerService,
	}
}

func (uh *FollowerHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.Follower
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = uh.followerService.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
