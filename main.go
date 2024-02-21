package main

import (
	"fmt"
	"net/http"
	"restrpcrouter/openapi"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// router := http.ServeMux()

	router.HandleFunc("/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)["path"]
		fmt.Println("Received path:", path)

		// Assuming openapi.GetSpec is a function that returns something
		operation, err := openapi.LookupSpec("/"+path, &r.Method)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		fmt.Print("Calling gRPC method: ", *operation)
	})

	fmt.Println("Starting server...")
	http.ListenAndServe(":8000", router)
}
