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
	// Chama o novo método do serviço que retorna apenas slots livres
	availableSlots, err := c.service.GetAvailableSchedules()
	if err != nil {
		jsonError(w, "Falha ao consultar o banco de dados", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(availableSlots)
}

func (c *ScheduleController) ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Schedule // Usa a struct do banco
	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		jsonError(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	if order.IDTime == "" || order.CustomerName == "" {
		jsonError(w, "id_horario e nome_cliente sao obrigatorios", http.StatusBadRequest)
		return
	}

	fmt.Printf("Request for schedule received %s of the customer %s\n", order.IDTime, order.CustomerName)

	res, err := c.service.CreateSchedule(order)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "horario indisponivel" {
			// 409 Conflict é o status ideal para "recurso já existe"
			jsonError(w, "Este horario ja esta reservado", http.StatusConflict)
		} else if errMsg == "id_horario invalido" {
			// 400 Bad Request para entrada inválida
			jsonError(w, "O id_horario enviado nao e um slot valido (ex: 'slot_05')", http.StatusBadRequest)
		} else {
			// Erro genérico do servidor
			jsonError(w, "Falha ao agendar no banco de dados", http.StatusInternalServerError)
		}
		return
	}

	fmt.Printf("Inserted into DB with ID: %v\n", res.InsertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensagem":    "Agendamento realizado com sucesso!",
		"id_inserido": res.InsertedID,
	})
}
