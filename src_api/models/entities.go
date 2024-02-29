package models

type Transacao struct {
	Id			   int64  `json:"-"`
	IdCliente      int  `json:"-"`
	Valor      int64 `json:"valor"`
	Tipo string `json:"tipo"`
	Descricao string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}

type Cliente struct {
	Id  int  `json:"-"`
	Saldo      int64 `json:"total"`
	Limite      int64 `json:"limite"`
}