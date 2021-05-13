package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// params: {id} form data: name
func GetParameters(response http.ResponseWriter, request *http.Request) {
	//init form data
	error := request.ParseForm()

	//validate error
	if error != nil {
		log.Fatal(error)
	}

	//getting values
	params := mux.Vars(request)
	a := request.FormValue("name")
	var id string = params["id"]
	res := []string{a, id}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(res)
}

// params: {id} form data: name
func PostParameters(response http.ResponseWriter, request *http.Request) {

	//init form data
	error := request.ParseForm()

	//validate error
	if error != nil {
		log.Fatal(error)
	}

	//getting values
	params := mux.Vars(request)
	name := request.FormValue("name")
	var id string = params["id"]
	res := []string{name, id}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(res)
}

// params: {id} form data: name
func PutParameters(response http.ResponseWriter, request *http.Request) {
	//init form data
	error := request.ParseForm()

	//validate error
	if error != nil {
		log.Fatal(error)
	}

	//getting values
	params := mux.Vars(request)
	name := request.FormValue("name")
	var id string = params["id"]
	res := []string{name, id}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(res)
}

// params: {id} form data: name
func DeleteParameters(response http.ResponseWriter, request *http.Request) {
	//init form data
	error := request.ParseForm()

	//validate error
	if error != nil {
		log.Fatal(error)
	}

	//getting values
	params := mux.Vars(request)
	name := request.FormValue("name")
	var id string = params["id"]
	res := []string{name, id}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(res)
}
