package http

import (
	"net/http"

	"azyk/internal/middleware"
	"azyk/internal/usecase"
	"azyk/internal/usecase/auth"

	"github.com/gorilla/mux"
)

// Router объединяет роутинг приложения.
type Router struct {
	mux         *mux.Router
	UserUsecase usecase.UserUsecase
	AuthUsecase auth.AuthUseCase

	// Обработчики для составов
	AminoHandler    usecase.AminoAcidCompositionService
	ChemicalHandler usecase.ChemicalCompositionService
	MineralHandler  usecase.MineralCompositionService
	FattyHandler    usecase.FattyAcidCompositionService
	VitaminHandler  usecase.VitaminCompositionService

	// Можно добавить и другие обработчики, например для продуктов
}

// Новый конструктор Router. Он принимает зависимости (usecase и обработчики)
// и регистрирует маршруты.
func NewRouter(
	userUsecase usecase.UserUsecase,
	authUsecase auth.AuthUseCase,
	aminoHandler usecase.AminoAcidCompositionService,
	chemicalHandler usecase.ChemicalCompositionService,
	mineralHandler usecase.MineralCompositionService,
	fattyHandler usecase.FattyAcidCompositionService,
	vitaminHandler usecase.VitaminCompositionService,
) *Router {
	r := &Router{
		mux:             mux.NewRouter(),
		UserUsecase:     userUsecase,
		AuthUsecase:     authUsecase,
		AminoHandler:    aminoHandler,
		ChemicalHandler: chemicalHandler,
		MineralHandler:  mineralHandler,
		FattyHandler:    fattyHandler,
		VitaminHandler:  vitaminHandler,
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

	// Маршруты для пользователей
	protected.HandleFunc("/users/logout", r.Logout).Methods("POST")
	protected.HandleFunc("/users/{id:[0-9]+}", r.GetUser).Methods("GET")
	protected.HandleFunc("/users/{id:[0-9]+}", r.UpdateUser).Methods("PUT")
	protected.HandleFunc("/users/{id:[0-9]+}", r.DeleteUser).Methods("DELETE")

	// Маршруты для состава аминокислот
	protected.HandleFunc("/compositions/amino", r.AminoHandler.CreateComposition).Methods("POST")
	protected.HandleFunc("/compositions/amino", r.AminoHandler.ListCompositionsByProduct).Methods("GET")
	protected.HandleFunc("/compositions/amino/{id:[0-9]+}", r.AminoHandler.GetCompositionByID).Methods("GET")
	protected.HandleFunc("/compositions/amino/{id:[0-9]+}", r.AminoHandler.UpdateComposition).Methods("PUT")
	protected.HandleFunc("/compositions/amino/{id:[0-9]+}", r.AminoHandler.DeleteComposition).Methods("DELETE")

	// Маршруты для химического состава
	protected.HandleFunc("/compositions/chemical", r.ChemicalCreateComposition).Methods("POST")
	protected.HandleFunc("/compositions/chemical", r.ChemicalListCompositionsByProduct).Methods("GET")
	protected.HandleFunc("/compositions/chemical/{id:[0-9]+}", r.ChemicalGetCompositionByID).Methods("GET")
	protected.HandleFunc("/compositions/chemical/{id:[0-9]+}", r.ChemicalUpdateComposition).Methods("PUT")
	protected.HandleFunc("/compositions/chemical/{id:[0-9]+}", r.ChemicalDeleteComposition).Methods("DELETE")

	// Маршруты для минерального состава
	protected.HandleFunc("/compositions/mineral", r.MineralCreateComposition).Methods("POST")
	protected.HandleFunc("/compositions/mineral", r.MineralListCompositionsByProduct).Methods("GET")
	protected.HandleFunc("/compositions/mineral/{id:[0-9]+}", r.MineralGetCompositionByID).Methods("GET")
	protected.HandleFunc("/compositions/mineral/{id:[0-9]+}", r.MineralUpdateComposition).Methods("PUT")
	protected.HandleFunc("/compositions/mineral/{id:[0-9]+}", r.MineralDeleteComposition).Methods("DELETE")

	// Маршруты для состава жирных кислот
	protected.HandleFunc("/compositions/fatty", r.FattyCreateComposition).Methods("POST")
	protected.HandleFunc("/compositions/fatty", r.FattyListCompositionsByProduct).Methods("GET")
	protected.HandleFunc("/compositions/fatty/{id:[0-9]+}", r.FattyGetCompositionByID).Methods("GET")
	protected.HandleFunc("/compositions/fatty/{id:[0-9]+}", r.FattyUpdateComposition).Methods("PUT")
	protected.HandleFunc("/compositions/fatty/{id:[0-9]+}", r.FattyDeleteComposition).Methods("DELETE")

	// Маршруты для витаминов
	protected.HandleFunc("/compositions/vitamin", r.VitaminCreateComposition).Methods("POST")
	protected.HandleFunc("/compositions/vitamin", r.VitaminListCompositionsByProduct).Methods("GET")
	protected.HandleFunc("/compositions/vitamin/{id:[0-9]+}", r.VitaminGetCompositionByID).Methods("GET")
	protected.HandleFunc("/compositions/vitamin/{id:[0-9]+}", r.VitaminUpdateComposition).Methods("PUT")
	protected.HandleFunc("/compositions/vitamin/{id:[0-9]+}", r.VitaminDeleteComposition).Methods("DELETE")
}

// ServeHTTP реализует интерфейс http.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
