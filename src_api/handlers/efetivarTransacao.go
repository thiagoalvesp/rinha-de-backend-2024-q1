package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/models"
)

func EfetivarTransacao(w http.ResponseWriter, r *http.Request) {
	var transacao models.Transacao

	Sid := chi.URLParam(r, "id")
	idCliente, err := strconv.Atoi(Sid)
	if err != nil {
		log.Printf("Erro ao fazer o parse do idCliente: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	
	err = json.NewDecoder(r.Body).Decode(&transacao)
	if err != nil {
		log.Printf("erro ao fazer decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	
	//cliente existe?
	_, err = models.BuscarClientePorId(idCliente)
	if err != nil {
		log.Printf("erro ao fazer decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	
	//transacao
	transacao.IdCliente = idCliente
	err = models.Efetivar(transacao)
	if err != nil {
		log.Printf("erro ao fazer decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	//consulto o saldo atual do cliente
	cliente, err := models.BuscarClientePorId(idCliente)
	if err != nil {
		log.Printf("erro ao fazer decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
	w.WriteHeader(http.StatusOK)
	
}
