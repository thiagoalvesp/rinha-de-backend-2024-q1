package main

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/handlers"
)

func main() {

	//pool de conexao pg
	dbConfig, err := pgxpool.ParseConfig("host=localhost port=5432 user=admin password=123 dbname=rinha sslmode=disable")
	if err!=nil {
		log.Fatal("Falha ao criar a configuracao, error: ", err)
	}

	dbConfig.MaxConns = int32(90)
	dbConfig.MinConns = int32(4)
	dbConfig.MaxConnLifetime = time.Hour
	dbConfig.MaxConnIdleTime = time.Minute * 30
	dbConfig.HealthCheckPeriod = time.Minute
	dbConfig.ConnConfig.ConnectTimeout = time.Second * 10

	// Create database connection
	connPool,err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err!=nil {
		log.Fatal("Erro durante a conexao com o banco")
	}
	defer connPool.Close()

	//router
	var p handlers.Pool
	p.ConnPool = connPool

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
