package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/configs"
	"log"

	_ "github.com/lib/pq"
)

func Config() (*pgxpool.Config) {
//	const defaultMaxConns = int32(4)
//	const defaultMinConns = int32(0)
//	const defaultMaxConnLifetime = time.Hour
//	const defaultMaxConnIdleTime = time.Minute * 30
//	const defaultHealthCheckPeriod = time.Minute
//	const defaultConnectTimeout = time.Second * 5

	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)


	dbConfig, err := pgxpool.ParseConfig(sc)
	if err!=nil {
		log.Fatal("Falha ao criar a configuracao, error: ", err)
	}

	return dbConfig
}


func OpenConnection() (*pgxpool.Pool, error) {

	// Create database connection
	connPool,err := pgxpool.NewWithConfig(context.Background(), Config())
	if err!=nil {
		log.Fatal("Erro durante a conexao com o banco")
	}

	connection, err := connPool.Acquire(context.Background())
	if err!=nil {
		log.Fatal("Erro durante aquisicao da conexao com o banco:", err)
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err!=nil{
		log.Fatal("Nao pode pingar o banco")
	}

	fmt.Println("Conectado no banco")

	return connPool, err
}
