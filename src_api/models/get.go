package models

import "github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/db"

func Get(idCliente int64) (transacao Transacao, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM transacoes WHERE idCliente=$1`, idCliente)

	err = row.Scan(&transacao.Id, &transacao.IdCliente, &transacao.Valor, &transacao.Descricao, &transacao.RealizadaEm)

	return
}
