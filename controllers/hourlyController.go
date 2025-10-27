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

type scheduleTime struct {
	IDTime       string `json:"id_time"`
	customerName string `json:"CustomerName"`
}

func (c *Controller) ConsultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Hello World Consult - Soon consulting the DB"})
}

func (c *Controller) ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var order scheduleTime
	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "JSON invalid"})
		return
	}

	fmt.Printf("Request for schedule received %s of the customer %s\n", order.IDTime, order.customerName)

	collection := c.Db.Collection("agendamentos")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, order)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Failed to schedule in database"})
		return
	}

	fmt.Printf("Inserted into DB with ID: %v\n", res.InsertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // StatusOK (200) ou StatusCreated (201)
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Appointment received successfully!"})
}
