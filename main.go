package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"consultancy_hours/controllers"
	"consultancy_hours/services"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not defined in .env")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Erro to the connect MongoDB: ", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Erro to the ping MongoDB: ", err)
	}

	fmt.Println("Connect to the MongoDB!")

	db := client.Database("agenda")

	scheduleService := services.NewScheduleService(db)
	appController := controllers.NewScheduleController(scheduleService)

	http.HandleFunc("/hours/available", appController.ConsultHandler)
	http.HandleFunc("/to schedule", appController.ScheduleHandler)

	fmt.Println("Iniciando servidor na porta 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error to the start server: ", err)
	}
}

//curl -X POST http://localhost:8080/toSchedule -H "Content-Type: application/json" -d '{"id_horario": "12:00", "nome_cliente": "Joao Gabriel"}'
//curl -X POST http://localhost:8080/toSchedule -H "Content-Type: application/json" -d '{"id_horario": "13:00", "nome_cliente": "Pedro Ivo"}'
//curl -X POST http://localhost:8080/toSchedule -H "Content-Type: application/json" -d '{"id_horario": "14:00", "nome_cliente": "Lais"}'

//curl http://localhost:8080/hours/available | jq .         ---> "sudo apt install jq"
