package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/configs"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(0)
	}

	r := chi.NewRouter()
	r.Post("/clientes/{id}/transacoes", handlers.EfetivarTransacao)
	r.Get("/clientes/{id}/extrato", handlers.Get)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(0)
	}

}
