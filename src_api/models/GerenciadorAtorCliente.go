package models

//import (
//	"fmt"
//)

type GerenciadorAtorCliente struct {
	Clientes map[int]*Cliente
}

func NovoGerenciadorAtorCliente() *GerenciadorAtorCliente {
	return &GerenciadorAtorCliente{Clientes: make(map[int]*Cliente)}
}

func (g *GerenciadorAtorCliente) RegistrarCliente(c Cliente) {
	g.Clientes[c.Id] = &c
	//fmt.Println("cliente", c.Id, "foi registrado no gerenciador")
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

func (c *Cliente) ProcessarMensagens() {
	for _, t := range c.Transacoes {
		//fmt.Println("cliente id:", t.IdCliente, "recebeu a transacao", t.Descricao)

		//saldo em memoria
		switch t.Tipo {
		case `c`:
			c.Creditar(t.Valor)
			break
		case `d`:
			c.Debitar(t.Valor)
			break
		}

		//registrar no banco
		//t.Efetivar()
	}
	//c.Transacoes = nil
}

func (c *Cliente) ReceberTransacao(t Transacao) {
	c.Transacoes = append(c.Transacoes, t)
}

