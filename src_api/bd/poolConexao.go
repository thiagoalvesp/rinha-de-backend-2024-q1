package bd

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
	"time"
)

var poolConexao *pgxpool.Pool
var once sync.Once

func RetornaPool() *pgxpool.Pool{
	once.Do(func() {
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
	

		poolConexao = connPool
	})
	return poolConexao
}