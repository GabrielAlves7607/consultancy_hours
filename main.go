package main

import (
	"fmt"
	"net/http" //
)

func main() {

	// http.HandleFunc --> used to define the route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World \n")

	})
	fmt.Println("Initing on port 8080")
	http.ListenAndServe(":8080", nil)

}
