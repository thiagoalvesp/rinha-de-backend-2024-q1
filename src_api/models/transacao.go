package models

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Efetivar(cliente Cliente, transacao Transacao, connPoll *pgxpool.Pool) (clienteRetorno Cliente, err error) {

	if transacao.DoTipoDebito() {
		if !cliente.PodeDebitar(transacao.Valor) {
			return cliente, errors.New("saldo insuficiente")
		}
	}

	switch transacao.Tipo {
	case `c`:
		cliente.Creditar(transacao.Valor)
		break
	case `d`:
		cliente.Debitar(transacao.Valor)
		break
	default:
		return cliente, errors.New("tipo de transacao invalido")
	}

	//Registrar no bd
//	_, err = connPoll.Exec(context.Background(), `INSERT INTO transacoes (idCliente, valor, tipo, descricao) VALUES ($1, $2, $3, $4) RETURNING id`, cliente.Id, transacao.Valor, transacao.Tipo, transacao.Descricao)
//	if err == nil {
//		if transacao.DoTipoDebito() {
//			transacao.Valor = transacao.Valor * -1
//		}
//		_, err = connPoll.Exec(context.Background(), `UPDATE clientes SET saldo= saldo + $1 WHERE id=$2`, transacao.Valor, cliente.Id)
//		//reverter
//		if err == nil {
//
//		}
//	}


	return cliente, err
}
