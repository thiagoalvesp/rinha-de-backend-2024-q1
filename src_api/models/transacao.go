package models

import (
	"context"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/bd"
	"log"
	"time"
)

func filter(transacoes []Transacao, predicate func(Transacao) bool) []Transacao {
	var result []Transacao
	for _, t := range transacoes {
		if predicate(t) {
			result = append(result, t)
		}
	}
	return result
}

type Transacao struct {
	Id          int64  `json:"-"`
	IdCliente   int    `json:"-"`
	Valor       int64  `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
	Efetivada   bool 	`json:"-"`
}

func (t *Transacao) DoTipoDebito() bool {
	return t.Tipo == "d"
}

func (t *Transacao) Efetivar () {

	////fazer bulk insert
	connPoll := bd.RetornaPool()
	_, err := connPoll.Exec(context.Background(), `INSERT INTO transacoes (idCliente, valor, tipo, descricao, realizada_em) VALUES ($1, $2, $3, $4, $5) RETURNING id`, t.IdCliente, t.Valor, t.Tipo, t.Descricao, t.RealizadaEm)
	if err != nil {
		log.Printf("erro ao inserir no banco de dados: %v", err)
	}
	if t.DoTipoDebito() {
		t.Valor = t.Valor * -1
	}
	_, err = connPoll.Exec(context.Background(), `UPDATE clientes SET saldo= saldo + $1 WHERE id=$2`, t.Valor, t.IdCliente)
	if err != nil {
		log.Printf("erro ao atualizar no banco de dados: %v", err)
	}
}