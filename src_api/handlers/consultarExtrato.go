package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (p ParamHandler) ConsultarExtrato(w http.ResponseWriter, r *http.Request) {

	Sid := chi.URLParam(r, "id")
	id, err := strconv.Atoi(Sid)

	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//cliente existe
	cliente, ok :=  p.GerenciadorAtorCliente.RetornaClienteAtorPorId(id)
	if !ok {
		log.Printf("cliente nao encontrado")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//carregar o extrato
	extrato, err := cliente.CarregarExtrato()
	if err != nil {
		log.Printf("Erro ao consultar o extrato: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(extrato)

}
