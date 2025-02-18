package http

import (
	"azyk/internal/middleware"
	"azyk/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	mux         *mux.Router
	UserUsecase usecase.UserUsecase
	AuthUsecase usecase.AuthUseCase
}

// Новый конструктор Router
func NewRouter(userUsecase usecase.UserUsecase, authUsecase usecase.AuthUseCase) *Router {
	r := &Router{
		mux:         mux.NewRouter(),
		UserUsecase: userUsecase,
		AuthUsecase: authUsecase,
	}
	r.registerRoutes() // Регистрируем маршруты
	return r
}

// Регистрация маршрутов
func (r *Router) registerRoutes() {
	// 🔹 Открытые маршруты (без авторизации)
	r.mux.HandleFunc("/users/register", r.RegisterUser).Methods("POST")
	r.mux.HandleFunc("/users/login", r.Login).Methods("POST")

	// 🔹 Защищенные маршруты (JWT-авторизация)
	protected := r.mux.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/users/logout", r.Logout).Methods("POST")
	protected.HandleFunc("/users/{id:[0-9]+}", r.GetUser).Methods("GET")
	protected.HandleFunc("/users/{id:[0-9]+}", r.UpdateUser).Methods("PUT")
	protected.HandleFunc("/users/{id:[0-9]+}", r.DeleteUser).Methods("DELETE")
}

// Реализация `http.Handler`
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
