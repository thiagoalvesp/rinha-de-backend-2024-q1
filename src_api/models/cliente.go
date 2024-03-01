package models

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func BuscarClientePorId(idCliente int, connPoll *pgxpool.Pool) (cliente Cliente, err error) {

	row := connPoll.QueryRow(context.Background(),`SELECT id, saldo, limite FROM clientes WHERE id=$1`, idCliente)
	err = row.Scan(&cliente.Id, &cliente.Saldo, &cliente.Limite)
	if err != nil {
		return 
	}
	
	return
}


func BuscarExtratoPorId(idCliente int, connPoll *pgxpool.Pool) (extrato Extrato, err error) {

	//carregar dados do cliente
	row := connPoll.QueryRow(context.Background(),`SELECT saldo, limite FROM clientes WHERE id=$1`, idCliente)
	err = row.Scan(&extrato.Saldo.Total, &extrato.Saldo.Limite)
	if err != nil {
		return
	}
	//carregar a data do extrato
	agora := time.Now()
	agoraFormatado := agora.Format("2006-01-02T15:04:05.999999Z")
	extrato.Saldo.DataExtrato = agoraFormatado

	log.Println("antes da query")
	//carregar dados das transacoes
	rows, err := connPoll.Query(context.Background(),"SELECT valor, tipo, descricao,  TO_CHAR(realizada_em, 'YYYY-MM-DD\"T\"HH24:MI:SS.US\"Z\"') realizada_em FROM transacoes WHERE idCliente = 1 ORDER BY id DESC LIMIT 10")
	if err != nil {
		return
	}
	defer rows.Close()

	///tem oportunidade aqui?
	for rows.Next() {
		var transacao Transacao
		err = rows.Scan(&transacao.Valor, &transacao.Tipo, &transacao.Descricao, &transacao.RealizadaEm)
		if err != nil {
			return
		}
		// Adicione o nome ao slice
		extrato.UltimasTransacoes = append(extrato.UltimasTransacoes, transacao)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return extrato, err
}