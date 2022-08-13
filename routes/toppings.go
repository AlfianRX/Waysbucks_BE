package routes

import (
	"back_end_waysbucks/handlers"
	"back_end_waysbucks/pkg/mysql"
	"back_end_waysbucks/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(userRepository)

	r.HandleFunc("/topping", h.FindToppings).Methods("GET")
	r.HandleFunc("/topping/{id}", h.GetTopping).Methods("GET")
	r.HandleFunc("/topping", h.CreateTopping).Methods("POST")
	r.HandleFunc("/topping/{id}", h.UpdateTopping).Methods("PATCH")
	r.HandleFunc("/topping/{id}", h.DeleteTopping).Methods("DELETE")

}
