package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/models"
)

type Pool struct {
	ConnPool *pgxpool.Pool
	GerenciadorAtorCliente *models.GerenciadorAtorCliente
}