package models

import (
	"fmt"
)

type Transacao struct {
	Id          int64  `json:"-"`
	IdCliente   int    `json:"-"`
	Valor       int64  `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}

func (t *Transacao) DoTipoDebito() bool {
	return t.Tipo == "d"
}

type GerenciadorAtorCliente struct {
	Clientes map[int]*Cliente
}

func NovoGerenciadorAtorCliente() *GerenciadorAtorCliente {
	return &GerenciadorAtorCliente{Clientes: make(map[int]*Cliente)}
}

func (g *GerenciadorAtorCliente) RegistrarCliente(c Cliente) {
	g.Clientes[c.Id] = &c
	fmt.Println("cliente", c.Id, "foi registrado no gerenciador")
}

func (g *GerenciadorAtorCliente) ReceberTransacao(t Transacao) {
	cliente, ok := g.Clientes[t.IdCliente]
	if ok {
		cliente.ReceberTransacao(t)
	}
}

func (g *GerenciadorAtorCliente) RetornaClienteAtorPorId(i int) (c *Cliente, ok bool) {
	cliente, ok := g.Clientes[i]
	return cliente, ok
}

type Cliente struct {
	Id             int         `json:"-"`
	Saldo          int64       `json:"total"`
	Limite         int64       `json:"limite"`
	CaixaMensagens []Transacao `json:"-"`
}

func (c *Cliente) ProcessarMensagens() {
	for _, t := range c.CaixaMensagens {
		fmt.Println("cliente id:", t.IdCliente, "recebeu a transacao", t.Descricao)

		switch t.Tipo {
		case `c`:
			c.Creditar(t.Valor)
			break
		case `d`:
			c.Debitar(t.Valor)
			break
		}

		//inserir no banco

	}
	c.CaixaMensagens = nil
}

func (c *Cliente) ReceberTransacao(t Transacao) {
	c.CaixaMensagens = append(c.CaixaMensagens, t)
}

//regras
func (c *Cliente) PodeDebitar(valor int64) bool {
	return ((c.Limite + c.Saldo) - valor) > 0
}

func (c *Cliente) Debitar(valor int64) {
	c.Saldo -= valor
}

func (c *Cliente) Creditar(valor int64) {
	c.Saldo += valor
}

type Saldo struct {
	Total       int64  `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int64  `json:"limite"`
}

type Extrato struct {
	Saldo             Saldo       `json:"saldo"`
	UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}
