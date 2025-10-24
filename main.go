package main

import (
	"consultancy_hours/controllers"
	"fmt"
	"net/http"
	//	"go.mongodb.org/mongo-driver/v2/mongo"
	//	"go.mongodb.org/mongo-driver/v2/mongo/options"
	//	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {

	//options.Client().ApplyURI("mongodb+srv://bagual:<db_password>@agenda.ipvs0d5.mongodb.net/?appName=agenda").SetServerAPIOptions(serverAPI)

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
