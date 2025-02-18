package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"azyk/internal/domain/models"
)

func (r *Router) RegisterUser(w http.ResponseWriter, req *http.Request) {
	var userReq models.CreateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&userReq); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := r.UserUsecase.RegisterUser(userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (r *Router) Login(w http.ResponseWriter, req *http.Request) {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		DeviceID string `json:"device_id"`
	}
	if err := json.NewDecoder(req.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, token, err := r.AuthUsecase.Login(loginReq.Email, loginReq.Password, loginReq.DeviceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp := map[string]interface{}{
		"user":  user,
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (r *Router) Logout(w http.ResponseWriter, req *http.Request) {
	var logoutReq struct {
		UserID int `json:"user_id"`
	}
	if err := json.NewDecoder(req.Body).Decode(&logoutReq); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := r.AuthUsecase.Logout(logoutReq.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

func (r *Router) GetUser(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	user, err := r.UserUsecase.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (r *Router) UpdateUser(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	var userReq models.UpdateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&userReq); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := r.UserUsecase.UpdateUser(id, userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func (r *Router) DeleteUser(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	err := r.UserUsecase.DeleteUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
