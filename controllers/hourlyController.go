package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	Db *mongo.Database
}

func NewController(db *mongo.Database) *Controller {
	return &Controller{Db: db}
}

type OrderMark struct {
	IDHorario   string `json:"id_horario"`
	NomeCliente string `json:"nome_cliente"`
}

func (c *Controller) ConsultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Hello World Consult - Em breve consultando o DB"})
}

func (c *Controller) ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var pedido OrderMark
	err := json.NewDecoder(r.Body).Decode(&pedido)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "JSON invalido"})
		return
	}

	fmt.Printf("Recebido pedido para o horario %s do cliente %s\n", pedido.IDHorario, pedido.NomeCliente)

	collection := c.Db.Collection("agendamentos")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, pedido)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Falha ao agendar no banco de dados"})
		return
	}

	fmt.Printf("Inserido no DB com ID: %v\n", res.InsertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // StatusOK (200) ou StatusCreated (201)
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Agendamento recebido com sucesso!"})
}
