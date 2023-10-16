package main

import (
	"fmt"
	"time"
)

type Pedido struct {
	ID          int
	Entrega     bool
	Produtos    []ProdutoPedido
	ValorTotal  float64
	CriadoEm    time.Time
	ExpedidoEm  time.Time
}

type ProdutoPedido struct {
	Produto  Produto
	Quantidade int
}

var Pedidos []Pedido
var proximoIDPedido = 1
const taxaEntrega = 10.0

// Adiciona um novo pedido à lista de pedidos
func adicionarPedido(entrega bool, produtosPedido []ProdutoPedido) {
	valor := calcularValorTotal(produtosPedido, entrega)
	pedido := Pedido{
		ID:          proximoIDPedido,
		Entrega:     entrega,
		Produtos:    produtosPedido,
		ValorTotal:  valor,
		CriadoEm:    time.Now(),
	}
	Pedidos = append(Pedidos, pedido)
	proximoIDPedido++
}

// Calcula o valor total do pedido considerando a taxa de entrega se aplicável
func calcularValorTotal(produtosPedido []ProdutoPedido, entrega bool) float64 {
	valor := 0.0
	for _, item := range produtosPedido {
		valor += item.Produto.Valor * float64(item.Quantidade)
	}
	if entrega {
		valor += taxaEntrega
	}
	return valor
}

// Expede o primeiro pedido da lista que ainda não foi expedido
func expedirPedido() *Pedido {
	for _, pedido := range Pedidos {
		if pedido.ExpedidoEm.IsZero() { // Verifica se o pedido ainda não foi expedido
			pedido.ExpedidoEm = time.Now()
			fmt.Printf("Pedido ID %d expedido!\n", pedido.ID)
			return &pedido
		}
	}
	fmt.Println("Todos os pedidos já foram expedidos.")
	return nil
}

// Exibe todos os pedidos em andamento (não expedidos)
func exibirPedidosEmAndamento() {
	for _, pedido := range Pedidos {
		if pedido.ExpedidoEm.IsZero() {
			fmt.Printf("Pedido ID: %d, Valor Total: R$%.2f, Criado em: %s\n", pedido.ID, pedido.ValorTotal, pedido.CriadoEm.Format("02/01/2006 15:04:05"))
		}
	}
}
