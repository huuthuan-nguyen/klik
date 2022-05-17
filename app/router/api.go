package router

import (
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/klik-dokter/app/handler"
	"github.com/huuthuan-nguyen/klik-dokter/app/middleware"
	"github.com/huuthuan-nguyen/klik-dokter/config"
	"net/http"
)

func SetAPIRoutes(router *mux.Router, config *config.Config, handler *handler.Handler) *mux.Router {
	apiRouter := router.PathPrefix("/api").Subrouter()

	// authentication
	apiRouter.HandleFunc("/register", handler.UserRegister).Methods(http.MethodPost, http.MethodOptions)
	apiRouter.HandleFunc("/auth/login", handler.UserLogin).Methods(http.MethodPost, http.MethodOptions)

	// products
	productRouter := apiRouter.PathPrefix("/").Subrouter()
	productRouter.HandleFunc("/products", handler.ProductStore).Methods(http.MethodPost, http.MethodOptions)
	productRouter.HandleFunc("/products/{id:[0-9]+}", handler.ProductUpdate).Methods(http.MethodPut, http.MethodOptions)
	productRouter.HandleFunc("/products/{id:[0-9]+}", handler.ProductDestroy).Methods(http.MethodDelete, http.MethodOptions)
	productRouter.HandleFunc("/products", handler.ProductIndex).Methods(http.MethodGet, http.MethodOptions)
	productRouter.HandleFunc("/products/{id:[0-9]+}", handler.ProductShow).Methods(http.MethodGet, http.MethodOptions)

	jwtAuthenticationMiddleware := middleware.JWTAuthenticateMiddleware(config, handler.GetDB())
	productRouter.Use(jwtAuthenticationMiddleware)

	return router
}
