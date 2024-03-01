package models

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

//retonar o cliente nessa transacao
func Efetivar(transacao Transacao, connPoll *pgxpool.Pool) (err error) {

	switch transacao.Tipo {
	case `c`:
		//Inserir Transacao
		err = EfetivarCredito(transacao, connPoll)
		break
	case `d`:
		err = EfetivarDebito(transacao, connPoll)
		break
	}

	return err
}

func EfetivarDebito(transacao Transacao, connPoll *pgxpool.Pool) error {

	// Inicie a transação
	tx, err := connPoll.Begin(context.Background())

	// Validacao se o cliente tem saldo + limite para efetuar a transacao
	row := tx.QueryRow(context.Background(),`select (((limite + saldo) - $1) >= 0) podeEfetivar from clientes  where id=$2`, transacao.Valor, transacao.IdCliente)
	podeEfetivar := false
	row.Scan(&podeEfetivar)
	if podeEfetivar {
		//Inserir Transacao
		_, err = tx.Exec(context.Background(),`INSERT INTO transacoes (idCliente, valor, tipo, descricao) VALUES ($1, $2, $3, $4) RETURNING id`, transacao.IdCliente, transacao.Valor, transacao.Tipo, transacao.Descricao)
		//Atualizar o saldo
		if err == nil {
			_, err = tx.Exec(context.Background(),`UPDATE clientes SET saldo= saldo + $1 WHERE id=$2`, transacao.Valor*-1, transacao.IdCliente)
		}
	} else {
		return errors.New("saldo insuficiente")
	}

	if err != nil {
		tx.Rollback(context.Background())
	} else {
		tx.Commit(context.Background())
	}
	
	return err
}

func EfetivarCredito(transacao Transacao, connPoll *pgxpool.Pool) error {


	// Inicie a transação
	tx, err := connPoll.Begin(context.Background())
	
	_, err = tx.Exec(context.Background(),`INSERT INTO transacoes (idCliente, valor, tipo, descricao) VALUES ($1, $2, $3, $4) RETURNING id`, transacao.IdCliente, transacao.Valor, transacao.Tipo, transacao.Descricao)
	if err == nil {
		_, err = tx.Exec(context.Background(),`UPDATE clientes SET saldo= saldo + $1 WHERE id=$2`, transacao.Valor, transacao.IdCliente)
	}

	if err != nil {
		tx.Rollback(context.Background())
	} else {
		tx.Commit(context.Background())
	}

	return err
}
