package main

import (
	"fmt"
	"net/http"
	//"github.com/labstack/echo/v4" 	--> echo
)

func consultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Consult")
}

func ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Schedule")
}

func main() {

	http.HandleFunc("/horarios/disponiveis", consultHandler)
	http.HandleFunc("/agendar", ScheduleHandler)

	fmt.Println("Initing on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error strating server ", err)
	}

}

// curl http://localhost:8080/horarios/disponiveis 	--> consultHandler
// curl http://localhost:8080/agendar				--> schedule
