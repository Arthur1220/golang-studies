package main

import (
	"POO/FINAL/clientes"
	"POO/FINAL/contas"
	"fmt"
)

func main() {
	// Criação de Conta Corrente
	contaDeTuca := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Tuca Developer",
			CPF:       "123.456.789-00",
			Profissao: "Desenvolvedor Go",
		},
		NumeroAgencia: 1234,
		NumeroConta:   556677,
	}
	// Depósito inicial
	contaDeTuca.Depositar(1000)

	// Criação de Conta Poupança
	contaDaMaria := contas.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "Maria Oliveira",
			CPF:       "987.654.321-00",
			Profissao: "Arquiteta",
		},
		NumeroAgencia: 1234,
		NumeroConta:   889900,
		Operacao:      13,
	}
	contaDaMaria.Depositar(500)

	// 1. Visualização inicial
	fmt.Println("--- Saldos Iniciais ---")
	exibirSaldo(&contaDeTuca)
	exibirSaldo(&contaDaMaria)

	// 2. Testando Transferência entre Tipos Diferentes (Corrente -> Poupança)
	// Isso só é possível porque alteramos a assinatura do método Transferencia para receber a Interface
	fmt.Println("\n--- Transferindo 100 reais da Tuca(CC) para Maria(CP) ---")
	status := contaDeTuca.Transferencia(100, &contaDaMaria)
	fmt.Println("Status da transferência:", status)

	// 3. Pagamento de Boletos (Polimorfismo)
	fmt.Println("\n--- Pagamento de Boletos ---")
	pagarBoleto(&contaDeTuca, 400)  // Paga com Corrente
	pagarBoleto(&contaDaMaria, 100) // Paga com Poupança

	// 4. Visualização Final
	fmt.Println("\n--- Saldos Finais ---")
	exibirSaldo(&contaDeTuca)
	exibirSaldo(&contaDaMaria)
}

// pagarBoleto aceita qualquer coisa que implemente a interface Conta (definida no pacote contas)
// Isso remove a necessidade de ter uma interface local chamada 'verificarConta'
func pagarBoleto(conta contas.Conta, valor float64) {
	if conta.Sacar(valor) {
		fmt.Printf("Boleto de R$ %.2f pago com sucesso.\n", valor)
	} else {
		fmt.Printf("Falha ao pagar boleto: saldo insuficiente.\n")
	}
}

// exibirSaldo é uma função auxiliar genérica
func exibirSaldo(conta contas.Conta) {
	// Como 'Conta' é uma interface, não temos acesso direto a campos como 'Titular' aqui,
	// apenas aos métodos definidos na interface (GetSaldo).
	fmt.Printf("Saldo atual: R$ %.2f\n", conta.GetSaldo())
}
