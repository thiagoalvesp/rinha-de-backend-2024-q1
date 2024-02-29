package models

import (
	"errors"
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/db"
)

func Insert(transacao Transacao) (id int64, err error) {

	switch transacao.Tipo {
	case `c`:
		//Inserir Transacao
		err = EfetivarCredito(transacao, id)
		break
	case `d`:
		err = EfetivarDebito(transacao, id)
		break
	}

	return id, err
}

func EfetivarDebito(transacao Transacao, id int64) error {

	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Inicie a transação
	tx, err := conn.Begin()

	// Validacao se o cliente tem saldo + limite para efetuar a transacao
	row := tx.QueryRow(`select (((limite + saldo) - $1) >= 0) podeEfetivar from clientes  where id=$2`, transacao.Valor, transacao.IdCliente)
	podeEfetivar := false
	row.Scan(&podeEfetivar)
	if podeEfetivar {
		//Inserir Transacao
		query := `INSERT INTO transacoes (idCliente, valor, tipo, descricao) VALUES ($1, $2, $3, $4) RETURNING id`
		err = tx.QueryRow(query, transacao.IdCliente, transacao.Valor, transacao.Tipo, transacao.Descricao).Scan(&id)
		//Atualizar o saldo
		if err == nil {
			_, err = tx.Exec(`UPDATE clientes SET saldo= saldo + $1 WHERE id=$2`, transacao.Valor*-1, transacao.IdCliente)
		}
	} else {
		return errors.New("saldo insuficiente")
	}

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}

func EfetivarCredito(transacao Transacao, id int64) error {

	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Inicie a transação
	tx, err := conn.Begin()

	query := `INSERT INTO transacoes (idCliente, valor, tipo, descricao) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(query, transacao.IdCliente, transacao.Valor, transacao.Tipo, transacao.Descricao).Scan(&id)
	if err == nil {
		_, err = tx.Exec(`UPDATE clientes SET saldo= saldo + $1 WHERE id=$2`, transacao.Valor, transacao.IdCliente)
	}

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}
