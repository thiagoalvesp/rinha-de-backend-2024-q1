package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/bd"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/models"
	"log"
	"net/http"


	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/handlers"
)


func main() {

	//Init do pool
	connPool := bd.RetornaPool()
	defer connPool.Close()

	//configuracoes do atores
	clientes, err :=  models.BuscarTodosClientes()

	gerenciadorAtorCliente := models.NovoGerenciadorAtorCliente()
	for _, c := range clientes {
		gerenciadorAtorCliente.RegistrarCliente(c)
	}

	//router
	var p handlers.ParamHandler
	p.GerenciadorAtorCliente = gerenciadorAtorCliente

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/clientes/{id}/transacoes", p.EfetivarTransacao )
	r.Get("/clientes/{id}/extrato", p.ConsultarExtrato  )

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Servido ouvindo na porta 8080")

}
