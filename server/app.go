<<<<<<< HEAD
package server

import (
	"apigo/routers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func App() {
	myrouters := routers.GetRouters()

	server := &http.Server{
		Addr:           ":3000",
		Handler:        myrouters,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server starting successfully")
	fmt.Println("http://localhost:3000")
	log.Fatal(server.ListenAndServe())

}
=======
package server

import (
	"apigo/routers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func App() {
	myrouters := routers.GetRouters()

	server := &http.Server{
		Addr:           ":3000",
		Handler:        myrouters,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server starting successfully")
	fmt.Println("http://localhost:3000")
	log.Fatal(server.ListenAndServe())

}
>>>>>>> 5de4841c7bfb23a45949f3a77640c1bd0421895a
