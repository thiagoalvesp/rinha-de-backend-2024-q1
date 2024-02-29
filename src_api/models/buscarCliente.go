package models

import "github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/db"

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
