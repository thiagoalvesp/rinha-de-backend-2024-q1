package handlers

import "github.com/jackc/pgx/v5/pgxpool"

type Pool struct {
	ConnPool *pgxpool.Pool
}