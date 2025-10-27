package controllers

import (
	"consultancy_hours/models"
	"consultancy_hours/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type ScheduleController struct {
	service *services.ScheduleService
}

func NewScheduleController(s *services.ScheduleService) *ScheduleController {
	return &ScheduleController{service: s}
}

func jsonError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"erro": message})
}

func (c *ScheduleController) ConsultHandler(w http.ResponseWriter, r *http.Request) {
	availableSlots, err := c.service.GetAvailableSchedules()
	if err != nil {
		jsonError(w, "Failed to query database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(availableSlots)
}

func (c *ScheduleController) ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Schedule
	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		jsonError(w, "JSON invalid", http.StatusBadRequest)
		return
	}

	if order.IDTime == "" || order.CustomerName == "" {
		jsonError(w, "id_horario and nome_cliente are mandatory", http.StatusBadRequest)
		return
	}

	fmt.Printf("Request for schedule received %s of the customer %s\n", order.IDTime, order.CustomerName)

	res, err := c.service.CreateSchedule(order)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "horario indisponivel" {
			// 409 Conflict is the ideal status for "resource already exists"
			jsonError(w, "This time is already booked", http.StatusConflict)
		} else if errMsg == "id_horario invalid" {
			// 400 Bad Request para entrada inválida
			jsonError(w, "The id_horario sent is not a valid slot", http.StatusBadRequest)
		} else {
			// Generic server error
			jsonError(w, "Failed to schedule in database", http.StatusInternalServerError)
		}
		return
	}

	fmt.Printf("Inserted into DB with ID: %v\n", res.InsertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensagem":    "Shedule completed successfully!",
		"id_inserido": res.InsertedID,
	})
}
