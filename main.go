package main

import (
	"back_end_waysbucks/database"
	"back_end_waysbucks/pkg/mysql"
	"back_end_waysbucks/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("Server Running")
	http.ListenAndServe("localhost:5000", r)

}
