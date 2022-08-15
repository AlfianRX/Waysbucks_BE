package routes

import (
	"back_end_waysbucks/handlers"
	"back_end_waysbucks/pkg/middleware"
	"back_end_waysbucks/pkg/mysql"
	"back_end_waysbucks/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	r.HandleFunc("/product", h.FindProducts).Methods("GET")
	r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/product", middleware.UploadFile(h.CreateProduct)).Methods("POST")
	r.HandleFunc("/product/{id}", h.UpdateProduct).Methods("PATCH")
	r.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")

}
