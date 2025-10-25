package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"consultancy_hours/controllers"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI não definida no .env")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Erro ao conectar ao MongoDB: ", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Erro ao pingar o MongoDB: ", err)
	}

	fmt.Println("Conectado ao MongoDB!")

	db := client.Database("agenda")

	appController := controllers.NewController(db)

	http.HandleFunc("/horarios/disponiveis", appController.ConsultHandler)
	http.HandleFunc("/agendar", appController.ScheduleHandler)

	fmt.Println("Iniciando servidor na porta 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar servidor: ", err)
	}
}

//curl -X POST http://localhost:8080/agendar \
//     -H "Content-Type: application/json" \
//     -d '{"id_horario": "123456", "nome_cliente": "Joao Gabriel"}'
