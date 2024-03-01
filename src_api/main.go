package main

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/db"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/configs"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	//abrindo o pool de conexo
	var connPool *pgxpool.Pool
	connPool, err = db.OpenConnection()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer connPool.Close()

	var p handlers.Pool
	p.ConnPool = connPool

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/clientes/{id}/transacoes", p.EfetivarTransacao )
	r.Get("/clientes/{id}/extrato", p.ConsultarExtrato  )

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}
