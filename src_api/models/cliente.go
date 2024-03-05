package models

import (
	"context"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/bd"
	"sort"
	"time"
)

type Saldo struct {
	Total       int64  `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int64  `json:"limite"`
}

type Extrato struct {
	Saldo             Saldo       `json:"saldo"`
	UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}

type Cliente struct {
	Id             int         `json:"-"`
	Saldo          int64       `json:"total"`
	Limite         int64       `json:"limite"`
	Transacoes []Transacao `json:"-"`
}

func (c *Cliente) PodeDebitar(valor int64) bool {
	return ((c.Limite + c.Saldo) - valor) >= 0
}

func (c *Cliente) Debitar(valor int64) {
	c.Saldo -= valor
}

func (c *Cliente) Creditar(valor int64) {
	c.Saldo += valor
}


func (c *Cliente) CarregarExtrato() (extrato Extrato, err error) {

	//carregar dados do cliente no extrato
	extrato.Saldo.Total = c.Saldo
	extrato.Saldo.Limite = c.Limite

	//carregar a data do extrato
	layoutData := "2006-01-02T15:04:05.999999Z"
	agora := time.Now()
	agoraFormatado := agora.Format(layoutData)
	extrato.Saldo.DataExtrato = agoraFormatado


	//carregar dados das transacoes
	//ordernar por data decr
	less := func(i, j int) bool {
		return c.Transacoes[i].RealizadaEm.After(c.Transacoes[j].RealizadaEm)
	}
	sort.Slice(c.Transacoes, less)

	if c.Transacoes != nil{
		take := min(len(c.Transacoes), 10)
		for _, t := range  c.Transacoes[:take] {
			extrato.UltimasTransacoes = append(extrato.UltimasTransacoes, t)
		}
	}

	return extrato, err
}

func BuscarTodosClientes() (clientes []Cliente, err error) {

	connPoll := bd.RetornaPool()

	rows, err := connPoll.Query(context.Background(),`SELECT id, saldo, limite FROM clientes ORDER BY id`)
	if err != nil {
		return
	}
	defer rows.Close()

	///tem oportunidade aqui?
	for rows.Next() {
		var c Cliente
		err = rows.Scan(&c.Id, &c.Saldo, &c.Limite)
		if err != nil {
			return
		}
		clientes = append(clientes, c)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return clientes, err
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}