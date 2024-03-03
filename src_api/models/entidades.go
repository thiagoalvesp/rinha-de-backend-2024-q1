package models

type Transacao struct {
	Id			   int64  `json:"-"`
	IdCliente      int  `json:"-"`
	Valor      int64 `json:"valor"`
	Tipo string `json:"tipo"`
	Descricao string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}

func (t *Transacao) DoTipoDebito() bool  {
	return t.Tipo == "d"
}

type Cliente struct {
	Id  int  `json:"-"`
	Saldo      int64 `json:"total"`
	Limite      int64 `json:"limite"`
}

func (c *Cliente) PodeDebitar(valor int64) bool {
	return 	((c.Limite + c.Saldo) - valor)  > 0
}

func (c *Cliente) Debitar(valor int64)  {
		c.Saldo -= valor
}

func (c *Cliente) Creditar(valor int64)  {
	c.Saldo += valor
}

type Saldo struct {
	Total      int64 `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int64 `json:"limite"`
}

type Extrato struct {
	Saldo Saldo `json:"saldo"`
	UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}

