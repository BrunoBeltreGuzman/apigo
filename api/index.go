<<<<<<< HEAD
package api

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, response *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello World!!")
}
=======
package api

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, response *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello World!!")
}
>>>>>>> 5de4841c7bfb23a45949f3a77640c1bd0421895a
