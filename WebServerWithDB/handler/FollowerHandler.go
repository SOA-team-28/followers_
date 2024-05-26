package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowerHandler struct {
	service *service.FollowerService
}

func NewFollowerHandler(driver neo4j.Driver) *FollowerHandler {
	followerService := service.NewFollowerService(driver)
	return &FollowerHandler{
		service: followerService,
	}
}

func (h *FollowerHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/follow", h.CreateFollowerHandler).Methods("POST")
	router.HandleFunc("/getById/{id}", h.GetById).Methods("GET")
	router.HandleFunc("/update/{existingUserID}/{newFollowerID}", h.UpdateFollower).Methods("PUT")

}

func (uh *FollowerHandler) CreateFollowerHandler(w http.ResponseWriter, r *http.Request) {
	var user model.Follower
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = uh.service.CreateFollower(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uh *FollowerHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// Uzmi ID Followera iz URL-a
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid follower ID", http.StatusBadRequest)
		return
	}

	// Pozovi servis za dobavljanje Followera
	follower, err := uh.service.GetById(id)
	if err != nil {
		http.Error(w, "Failed to get follower", http.StatusInternalServerError)
		return
	}

	// Serijalizuj Followera u JSON format i pošalji kao odgovor
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(follower)
	if err != nil {
		http.Error(w, "Failed to encode follower data", http.StatusInternalServerError)
		return
	}
}
func (uh *FollowerHandler) UpdateFollower(w http.ResponseWriter, r *http.Request) {
	// Uzmi ID postojećeg pratioca i ID novog pratioca iz URL-a
	vars := mux.Vars(r)
	existingUserID, err := strconv.Atoi(vars["existingUserID"])
	if err != nil {
		http.Error(w, "Invalid existing user ID", http.StatusBadRequest)
		return
	}
	newFollowerID, err := strconv.Atoi(vars["newFollowerID"])
	if err != nil {
		http.Error(w, "Invalid new follower ID", http.StatusBadRequest)
		return
	}

	// Pozovi servis za ažuriranje objekta pratioca
	err = uh.service.UpdateFollower(existingUserID, newFollowerID)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
