package main

import (
  "bufio"
  "os"
	"fmt"
	"strings"
)

func main() {
	for {
		exibirMenu()
		escolha := obterEscolhaUsuario()
		tratarEscolha(escolha)
	}
}

func exibirMenu() {
	fmt.Println("======= McRonald’s =======")
	fmt.Println("1. Cadastrar produto")
	fmt.Println("2. Remover produto")
	fmt.Println("3. Buscar produto por ID")
	fmt.Println("4. Exibir todos os produtos")
	fmt.Println("5. Adicionar pedido")
	fmt.Println("6. Expedir pedido")
	fmt.Println("7. Exibir pedidos em andamento")
	fmt.Println("8. Exibir métricas")
	fmt.Println("9. Sair")
	fmt.Println("==========================")
	fmt.Println("Digite sua opção:")
}

func obterEscolhaUsuario() int {
	var escolha int
	fmt.Scan(&escolha)
	return escolha
}

func tratarEscolha(escolha int) {
	switch escolha {
	case 1:
		cadastrarProduto()
	case 2:
		interfaceRemoverProduto()
	case 3:
		buscarProdutoPorIDInterface()
	case 4:
		exibirProdutos()
	case 5:
		adicionarPedidoInterface()
	case 6:
		expedirPedidoInterface()
	case 7:
		exibirPedidosEmAndamento()
	case 8:
		exibirMetricas()
	case 9:
		fmt.Println("Obrigado por usar o sistema McRonald’s!")
		exit()
	default:
		fmt.Println("Opção inválida. Tente novamente.")
	}
}

func cadastrarProduto() {
	var nome, descricao string
	var valor float64

	fmt.Println("Digite o nome do produto:")
	fmt.Scanln(&nome)
	nome += strings.TrimSpace(readRemainingLine())

	fmt.Println("Digite a descrição do produto:")
	fmt.Scanln(&descricao)
	descricao += strings.TrimSpace(readRemainingLine())

	fmt.Println("Digite o valor do produto:")
	fmt.Scan(&valor)

	adicionarProduto(nome, descricao, valor)
	fmt.Println("Produto cadastrado com sucesso!")
}

func interfaceRemoverProduto() {
	var id int
	fmt.Println("Digite o ID do produto a ser removido:")
	fmt.Scan(&id)
	removerProduto(id)  // chamando a função do arquivo produto.go
	fmt.Println("Produto removido com sucesso!")
}

func buscarProdutoPorIDInterface() {
	var id int
	fmt.Println("Digite o ID do produto a ser buscado:")
	fmt.Scan(&id)
	produto := buscarProdutoPorID(id)
	if produto != nil {
		fmt.Printf("Produto encontrado: Nome: %s, Descrição: %s, Valor: R$%.2f\n", produto.Nome, produto.Descricao, produto.Valor)
	} else {
		fmt.Println("Produto não encontrado.")
	}
}

func adicionarPedidoInterface() {
	// Para simplificar, vou considerar apenas a adição de um produto por pedido aqui.
	// Em um sistema real, um loop seria necessário para adicionar múltiplos produtos a um pedido.

	var idProduto, quantidade int
	var entrega bool
	var opcaoEntrega string

	fmt.Println("Digite o ID do produto:")
	fmt.Scan(&idProduto)
	produto := buscarProdutoPorID(idProduto)
	if produto == nil {
		fmt.Println("Produto não encontrado.")
		return
	}

	fmt.Println("Digite a quantidade do produto:")
	fmt.Scan(&quantidade)

	fmt.Println("É um pedido de entrega? (s/n)")
	fmt.Scan(&opcaoEntrega)
	if strings.ToLower(opcaoEntrega) == "s" {
		entrega = true
	}

	adicionarPedido(entrega, []ProdutoPedido{
		{
			Produto:   *produto,
			Quantidade: quantidade,
		},
	})
	fmt.Println("Pedido adicionado com sucesso!")
}

func expedirPedidoInterface() {
	pedidoExpedido := expedirPedido()
	if pedidoExpedido != nil {
		fmt.Printf("Pedido ID %d foi expedido com sucesso!\n", pedidoExpedido.ID)
	}
}

// Função utilitária para ler o restante da linha após um Scanln
func readRemainingLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func exit() {
	os.Exit(0)
}
