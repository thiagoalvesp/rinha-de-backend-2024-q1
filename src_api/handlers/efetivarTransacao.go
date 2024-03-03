package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/models"
)

func (p Pool) EfetivarTransacao(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	//cliente existe
	cliente, err := models.BuscarClientePorId(idCliente, p.ConnPool)
	if err != nil {
		log.Printf("erro ao buscar o cliente: %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//efetiva a transacao e retorna o saldo atualizado
	cliente, err = models.Efetivar(cliente, transacao, p.ConnPool)
	if err != nil {
		log.Printf("erro ao efetivar a transacao: %v", err)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	if err == nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cliente)
	}

}
