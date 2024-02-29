package models

import (
	"github.com/thiagoalvesp/rinha-de-backend-2024-q1/src_api/db"
)

func Update(id int64, novoSaldo int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE cliente SET saldo=$1 WHERE id=$2`, novoSaldo, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}