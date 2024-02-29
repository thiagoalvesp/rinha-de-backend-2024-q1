package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/models"
)

func ConsultarExtrato(w http.ResponseWriter, r *http.Request) {

	Sid := chi.URLParam(r, "id")
	id, err := strconv.Atoi(Sid)

	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//cliente existe?
	_, err = models.BuscarClientePorId(id)
	if err != nil {
		log.Printf("erro ao fazer decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//carregar o extrato
	extrato, err := models.BuscarExtratoPorId(id)
	if err != nil {
		log.Printf("Erro ao consultar o extrato: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(extrato)
	w.WriteHeader(http.StatusOK)
}
