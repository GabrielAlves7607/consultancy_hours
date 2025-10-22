package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderMark struct {
	IDHorario   string `json:"id_horario"`
	NomeCliente string `json:"nome_cliente"`
}

func ConsultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Hello World Consult"})
}

func ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var pedido OrderMark
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erro": "JSON invalido"})
		return
	}

	fmt.Printf("Recebido pedido para o horario %s do cliente %s\n", pedido.IDHorario, pedido.NomeCliente)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Agendamento recebido com sucesso!"})
}
