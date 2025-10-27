package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schedule struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IDTime       string             `json:"id_horario"   bson:"id_horario"`
	CustomerName string             `json:"nome_cliente" bson:"nome_cliente"`
}

type AvailableSlot struct {
	IDTime string `json:"id_horario"`
	Status string `json:"status"`
}
