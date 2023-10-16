package main

import (
	"fmt"
)

type Produto struct {
	ID          int
	Nome        string
	Descricao   string
	Valor       float64
}

var Produtos []Produto
var proximoIDProduto = 1

// Adiciona um novo produto à lista de produtos
func adicionarProduto(nome, descricao string, valor float64) {
	produto := Produto{
		ID:        proximoIDProduto,
		Nome:      nome,
		Descricao: descricao,
		Valor:     valor,
	}
	Produtos = append(Produtos, produto)
	proximoIDProduto++
}

// Remove um produto da lista de produtos pelo ID
func removerProduto(id int) {
	for indice, produto := range Produtos {
		if produto.ID == id {
			Produtos = append(Produtos[:indice], Produtos[indice+1:]...)
			return
		}
	}
}

// Busca um produto pela ID e retorna
func buscarProdutoPorID(id int) *Produto {
	for _, produto := range Produtos {
		if produto.ID == id {
			return &produto
		}
	}
	return nil
}

// Exibe todos os produtos cadastrados
func exibirProdutos() {
	if len(Produtos) == 0 {
		fmt.Println("Nenhum produto cadastrado.")
		return
	}
	for _, produto := range Produtos {
		fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Valor: R$%.2f\n", produto.ID, produto.Nome, produto.Descricao, produto.Valor)
	}
}

// Adiciona uma lista de produtos de uma vez
func adicionarProdutosEmLote(produtosLote []Produto) {
	for _, produto := range produtosLote {
		adicionarProduto(produto.Nome, produto.Descricao, produto.Valor)
	}
}

// Busca produtos pelo nome e retorna uma lista de produtos
func buscarProdutosPorNome(nome string) []Produto {
	var produtosEncontrados []Produto
	for _, produto := range Produtos {
		if nome == produto.Nome {
			produtosEncontrados = append(produtosEncontrados, produto)
		}
	}
	return produtosEncontrados
}
