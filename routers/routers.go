<<<<<<< HEAD
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

	routers.HandleFunc("/parameters/{id}", api.GetParameters).Methods("GET")
	routers.HandleFunc("/parameters/{id}", api.PostParameters).Methods("POST")
	routers.HandleFunc("/parameters/{id}", api.PutParameters).Methods("PUT")
	routers.HandleFunc("/parameters/{id}", api.DeleteParameters).Methods("DELETE")

	return routers
}
=======
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
>>>>>>> 5de4841c7bfb23a45949f3a77640c1bd0421895a
