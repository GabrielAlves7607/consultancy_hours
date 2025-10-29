package controllers

import (
	"consultancy_hours/models"
	"consultancy_hours/services"
	"encoding/json"
	"fmt"
	"net/http"

	"consultancy_hours/constants/constControllers"
)

type ScheduleController struct {
	service *services.ScheduleService
}

func NewScheduleController(s *services.ScheduleService) *ScheduleController {
	return &ScheduleController{service: s}
}

func jsonError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set(constControllers.HeaderContentType, constControllers.HeaderContentTypeJSON) // Constante
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{constControllers.JSONKeyError: message})
}

func (c *ScheduleController) ConsultHandler(w http.ResponseWriter, r *http.Request) {
	availableSlots, err := c.service.GetAvailableSchedules()
	if err != nil {
		jsonError(w, constControllers.MsgErrorFailedDBQuery, http.StatusInternalServerError)
		return
	}

	w.Header().Set(constControllers.HeaderContentType, constControllers.HeaderContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(availableSlots)
}

func (c *ScheduleController) ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Schedule
	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		jsonError(w, constControllers.MsgErrorInvalidJSON, http.StatusBadRequest)
		return
	}

	if order.IDTime == "" || order.CustomerName == "" {
		jsonError(w, constControllers.MsgErrorMissingFields, http.StatusBadRequest)
		return
	}

	fmt.Printf(constControllers.LogRequestReceived, order.IDTime, order.CustomerName)

	res, err := c.service.CreateSchedule(order)
	if err != nil {
		errMsg := err.Error()

		if errMsg == constControllers.ErrorScheduleUnavailable {
			jsonError(w, constControllers.MsgErrorScheduleUnavailable, http.StatusConflict)

		} else if errMsg == constControllers.ErrorTimeId {
			jsonError(w, constControllers.MsgErrorTimeId, http.StatusBadRequest)

		} else {
			jsonError(w, constControllers.MsgErrorDatabase, http.StatusInternalServerError)
		}
		return
	}

	fmt.Printf(constControllers.LogInsertSuccess, res.InsertedID)

	w.Header().Set(constControllers.HeaderContentType, constControllers.HeaderContentTypeJSON)
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(map[string]interface{}{
		constControllers.JSONKeyMessage:    constControllers.MsgScheduleSuccess,
		constControllers.JSONKeyInsertedID: res.InsertedID,
	})
}
