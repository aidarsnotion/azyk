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

	AminoHandler    *AminoAcidCompositionHandler
	ChemicalHandler *ChemicalCompositionHandler
	MineralHandler  *MineralCompositionHandler
	FattyHandler    *FattyAcidCompositionHandler
	VitaminHandler  *VitaminCompositionHandler
	ProductHandler  *ProductHandler
	CategoryHandler *CategoryHandler
}

// Новый конструктор Router. Он принимает зависимости (usecase и обработчики)
// и регистрирует маршруты.
func NewRouter(
	// юзкейсы
	userUsecase usecase.UserUsecase,
	authUsecase auth.AuthUseCase,

	// обработчики
	AminoHandler *AminoAcidCompositionHandler,
	ChemicalHandler *ChemicalCompositionHandler,
	MineralHandler *MineralCompositionHandler,
	FattyHandler *FattyAcidCompositionHandler,
	VitaminHandler *VitaminCompositionHandler,
	ProductHandler *ProductHandler,
	CategoryHandler *CategoryHandler,
) *Router {
	r := &Router{
		mux:         mux.NewRouter(),
		UserUsecase: userUsecase,
		AuthUsecase: authUsecase,

		AminoHandler:    AminoHandler,
		ChemicalHandler: ChemicalHandler,
		MineralHandler:  MineralHandler,
		FattyHandler:    FattyHandler,
		VitaminHandler:  VitaminHandler,
		ProductHandler:  ProductHandler,
		CategoryHandler: CategoryHandler,
	}
	r.registerRoutes()
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
	protected.HandleFunc("/compositions/chemical", r.ChemicalHandler.CreateComposition).Methods("POST")
	protected.HandleFunc("/compositions/chemical", r.ChemicalHandler.ListCompositionsByProduct).Methods("GET")
	protected.HandleFunc("/compositions/chemical/{id:[0-9]+}", r.ChemicalHandler.GetCompositionByID).Methods("GET")
	protected.HandleFunc("/compositions/chemical/{id:[0-9]+}", r.ChemicalHandler.UpdateComposition).Methods("PUT")
	protected.HandleFunc("/compositions/chemical/{id:[0-9]+}", r.ChemicalHandler.DeleteComposition).Methods("DELETE")

	// Маршруты для минерального состава
	protected.HandleFunc("/compositions/mineral", r.MineralHandler.CreateComposition).Methods("POST")
	protected.HandleFunc("/compositions/mineral", r.MineralHandler.ListCompositionsByProduct).Methods("GET")
	protected.HandleFunc("/compositions/mineral/{id:[0-9]+}", r.MineralHandler.GetCompositionByID).Methods("GET")
	protected.HandleFunc("/compositions/mineral/{id:[0-9]+}", r.MineralHandler.UpdateComposition).Methods("PUT")
	protected.HandleFunc("/compositions/mineral/{id:[0-9]+}", r.MineralHandler.DeleteComposition).Methods("DELETE")

	// Маршруты для состава жирных кислот
	protected.HandleFunc("/compositions/fatty", r.FattyHandler.CreateComposition).Methods("POST")
	protected.HandleFunc("/compositions/fatty", r.FattyHandler.ListCompositionsByProduct).Methods("GET")
	protected.HandleFunc("/compositions/fatty/{id:[0-9]+}", r.FattyHandler.GetCompositionByID).Methods("GET")
	protected.HandleFunc("/compositions/fatty/{id:[0-9]+}", r.FattyHandler.UpdateComposition).Methods("PUT")
	protected.HandleFunc("/compositions/fatty/{id:[0-9]+}", r.FattyHandler.DeleteComposition).Methods("DELETE")

	// Маршруты для витаминов
	protected.HandleFunc("/compositions/vitamin", r.VitaminHandler.CreateComposition).Methods("POST")
	protected.HandleFunc("/compositions/vitamin", r.VitaminHandler.ListCompositionsByProduct).Methods("GET")
	protected.HandleFunc("/compositions/vitamin/{id:[0-9]+}", r.VitaminHandler.GetCompositionByID).Methods("GET")
	protected.HandleFunc("/compositions/vitamin/{id:[0-9]+}", r.VitaminHandler.UpdateComposition).Methods("PUT")
	protected.HandleFunc("/compositions/vitamin/{id:[0-9]+}", r.VitaminHandler.DeleteComposition).Methods("DELETE")

	// Маршруты для продуктов
	protected.HandleFunc("/product", r.ProductHandler.CreateProduct).Methods("POST")
	protected.HandleFunc("/product", r.ProductHandler.ListProducts).Methods("GET")
	protected.HandleFunc("/product/{id:[0-9]+}", r.ProductHandler.GetProductByID).Methods("GET")
	protected.HandleFunc("/product/{id:[0-9]+}", r.ProductHandler.UpdateProduct).Methods("PUT")
	protected.HandleFunc("/product/{id:[0-9]+}", r.ProductHandler.DeleteProduct).Methods("DELETE")

	// Маршруты для категории
	protected.HandleFunc("/category", r.CategoryHandler.Create).Methods("POST")
	protected.HandleFunc("/category", r.CategoryHandler.List).Methods("GET")
	protected.HandleFunc("/category/{id:[0-9]+}", r.CategoryHandler.GetByID).Methods("GET")
	protected.HandleFunc("/category/{id:[0-9]+}", r.CategoryHandler.Update).Methods("PUT")
	protected.HandleFunc("/category/{id:[0-9]+}", r.CategoryHandler.Delete).Methods("DELETE")
}

// ServeHTTP реализует интерфейс http.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
