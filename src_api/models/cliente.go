package models

import (
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/db"
	"time"
)

func BuscarClientePorId(idCliente int) (cliente Cliente, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT id, saldo, limite FROM clientes WHERE id=$1`, idCliente)
	err = row.Scan(&cliente.Id, &cliente.Saldo, &cliente.Limite)
	if err != nil {
		return 
	}
	
	return
}


func BuscarExtratoPorId(idCliente int) (extrato Extrato, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	//carregar dados do cliente
	row := conn.QueryRow(`SELECT saldo, limite FROM clientes WHERE id=$1`, idCliente)
	err = row.Scan(&extrato.Saldo.Total, &extrato.Saldo.Limite)
	if err != nil {
		return
	}
	//carregar a data do extrato
	agora := time.Now()
	agoraFormatado := agora.Format("2006-01-02T15:04:05.999999Z")
	extrato.Saldo.DataExtrato = agoraFormatado

	//carregar dados das transacoes
	rows, err := conn.Query("SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE idCliente = 1 ORDER BY id DESC LIMIT 10")
	if err != nil {
		return
	}
	defer rows.Close()

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