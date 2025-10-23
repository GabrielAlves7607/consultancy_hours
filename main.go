package main

import (
	"consultancy_hours/controllers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/horarios/disponiveis", controllers.ConsultHandler)
	http.HandleFunc("/agendar", controllers.ScheduleHandler)

	fmt.Println("Initing on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error strating server ", err)
	}

}

//curl -X POST http://localhost:8080/agendar \
//     -H "Content-Type: application/json" \
//     -d '{"id_horario": "123456", "nome_cliente": "Joao Gabriel"}'
