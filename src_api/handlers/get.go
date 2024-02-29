package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/models"
)

func Get(w http.ResponseWriter, r *http.Request) {

	Sid := chi.URLParam(r, "id")
	id, err := strconv.Atoi(Sid)

	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//404 qndo nao retornar clientes
	
	transacao, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao buscar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transacao)
}
