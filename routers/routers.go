package routers

import (
	"apigo/api"

	"github.com/gorilla/mux"
)

func GetRouters() (r *mux.Router) {
	routers := mux.NewRouter()
	routers.HandleFunc("/", api.HomeHandler).Methods("GET")
	routers.HandleFunc("/users", api.FindAllUsers).Methods("GET")
	routers.HandleFunc("/users/{id}", api.FindByIdUsers).Methods("GET")
	routers.HandleFunc("/users/{id}", api.UpdateUsers).Methods("PUT")
	routers.HandleFunc("/users/{id}", api.DeleteUsers).Methods("DELETE")

	return routers
}
