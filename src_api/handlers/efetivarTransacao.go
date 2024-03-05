package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/models"
)

func (p ParamHandler) EfetivarTransacao(w http.ResponseWriter, r *http.Request) {
	var transacao models.Transacao

	Sid := chi.URLParam(r, "id")
	idCliente, err := strconv.Atoi(Sid)
	if err != nil {
		log.Printf("erro ao fazer o parse do idCliente: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//carregando a transacao com as info do request e a data de recebimento
	transacao.IdCliente = idCliente
	//carregar a data do extrato
	transacao.RealizadaEm = time.Now()

	err = json.NewDecoder(r.Body).Decode(&transacao)
	if err != nil {
		log.Printf("erro ao fazer decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	//cliente existe
	cliente, ok :=  p.GerenciadorAtorCliente.RetornaClienteAtorPorId(transacao.IdCliente)
	if !ok {
		log.Printf("cliente nao encontrado")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//pode processar
	if transacao.DoTipoDebito() {
		if !cliente.PodeDebitar(transacao.Valor) {
			log.Printf("saldo insuficiente")
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}
	}

	//envia transacao
	p.GerenciadorAtorCliente.ReceberTransacao(transacao)

	//processa do ator
	cliente.ProcessarMensagens()

	if err == nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cliente)
	}

}
