package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"azyk/internal/domain/models"
	"azyk/util/logger"
)

func (r *Router) RegisterUser(w http.ResponseWriter, req *http.Request) {
	var userReq models.CreateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&userReq); err != nil {
		logger.Log.WithError(err).Error("Не удалось декодировать тело запроса на регистрацию пользователя")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("email", userReq.Email).Info("Получен запрос на регистрацию пользователя")

	user, err := r.UserUsecase.RegisterUser(userReq)
	if err != nil {
		logger.Log.WithError(err).Error("Ошибка регистрации пользователя")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Log.WithField("userID", user.ID).Info("Пользователь успешно зарегистрирован")
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
		logger.Log.WithError(err).Error("Не удалось декодировать тело запроса на авторизацию")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("email", loginReq.Email).Info("Получен запрос на авторизацию")

	user, token, err := r.AuthUsecase.Login(loginReq.Email, loginReq.Password, loginReq.DeviceID)
	if err != nil {
		logger.Log.WithError(err).Warn("Ошибка авторизации пользователя")
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	logger.Log.WithField("email", loginReq.Email).Info("Пользователь успешно авторизовался")
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
		logger.Log.WithError(err).Error("Не удалось декодировать тело запроса на выход пользователя")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("userID", logoutReq.UserID).Info("Получен запрос на выход пользователя")
	err := r.AuthUsecase.Logout(logoutReq.UserID)
	if err != nil {
		logger.Log.WithError(err).Error("Ошибка выхода пользователя")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Log.WithField("userID", logoutReq.UserID).Info("Пользователь успешно вышел")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

func (r *Router) GetUser(w http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.WithError(err).Error("Неверный идентификатор пользователя в запросе на получение пользователя")
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("userID", id).Info("Получен запрос на получение пользователя")
	user, err := r.UserUsecase.GetUser(id)
	if err != nil {
		logger.Log.WithError(err).Warn("Пользователь не найден")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	logger.Log.WithField("userID", id).Info("Пользователь успешно получен")
	json.NewEncoder(w).Encode(user)
}

func (r *Router) UpdateUser(w http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.WithError(err).Error("Неверный идентификатор пользователя в запросе на обновление")
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	var userReq models.UpdateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&userReq); err != nil {
		logger.Log.WithError(err).Error("Не удалось декодировать тело запроса на обновление пользователя")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("userID", id).Info("Получен запрос на обновление пользователя")
	err = r.UserUsecase.UpdateUser(id, userReq)
	if err != nil {
		logger.Log.WithError(err).Error("Ошибка обновления пользователя")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Log.WithField("userID", id).Info("Пользователь успешно обновлен")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func (r *Router) DeleteUser(w http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.WithError(err).Error("Неверный идентификатор пользователя в запросе на удаление")
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("userID", id).Info("Получен запрос на удаление пользователя")
	err = r.UserUsecase.DeleteUser(id)
	if err != nil {
		logger.Log.WithError(err).Warn("Пользователь не найден для удаления")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	logger.Log.WithField("userID", id).Info("Пользователь успешно удален")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
