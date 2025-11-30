package main

import (
	"POO/05/clientes"
	"POO/05/contas"
	"fmt"
)

func main() {

	contaCorrente := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "João Silva",
			CPF:       "123.456.789-00",
			Profissao: "Engenheiro",
		},
		NumeroAgencia: 1234,
		NumeroConta:   567890,
	}

	contaPoupanca := contas.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "João Silva",
			CPF:       "123.456.789-00",
			Profissao: "Engenheiro",
		},
		NumeroAgencia: 1234,
		NumeroConta:   567890,
	}

	printContaCorrente(contaCorrente)
	printContaPoupanca(contaPoupanca)

	contaCorrente2 := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Maria Oliveira",
			CPF:       "987.654.321-00",
			Profissao: "Advogada",
		},
		NumeroAgencia: 4321,
		NumeroConta:   987654,
	}

	printContaCorrente(contaCorrente2)

	fmt.Println("\nOperações Bancárias:")

	sucesso := contaCorrente.Sacar(50.00)
	if sucesso {
		fmt.Printf("Saque realizado com sucesso! Novo saldo: %.2f\n", contaCorrente.GetSaldo())
	} else {
		fmt.Println("Saldo insuficiente para o saque.")
	}

	sucesso = contaCorrente.Depositar(150.00)
	if sucesso {
		fmt.Printf("Depósito realizado com sucesso! Novo saldo: %.2f\n", contaCorrente.GetSaldo())
	} else {
		fmt.Println("Valor de depósito inválido.")
	}

	sucesso = contaCorrente.Transferencia(75.00, &contaCorrente2)
	if sucesso {
		fmt.Printf("Transferência realizada com sucesso! Novo saldo da conta de %s: %.2f\n", contaCorrente.Titular, contaCorrente.GetSaldo())
		fmt.Printf("Novo saldo da conta de %s: %.2f\n", contaCorrente2.Titular, contaCorrente2.GetSaldo())
	} else {
		fmt.Println("Saldo insuficiente para a transferência.")
	}

	sucesso = contaPoupanca.Depositar(300.00)
	if sucesso {
		fmt.Printf("Depósito realizado com sucesso na poupança! Novo saldo: %.2f\n", contaPoupanca.GetSaldo())
	} else {
		fmt.Println("Valor de depósito inválido na poupança.")
	}

	pagarBoleto(&contaCorrente, 50.00)
	fmt.Printf("Após pagar boleto, novo saldo da conta corrente: %.2f\n", contaCorrente.GetSaldo())
	pagarBoleto(&contaPoupanca, 10.00)
	fmt.Printf("Após pagar boleto, novo saldo da conta poupança: %.2f\n", contaPoupanca.GetSaldo())

	fmt.Println("\nDados Atualizados das Contas:")

	printContaCorrente(contaCorrente)
	printContaPoupanca(contaPoupanca)
	printContaCorrente(contaCorrente2)
}

func printContaCorrente(conta contas.ContaCorrente) {
	fmt.Println("\nDados da Conta Bancária Corrente:")
	fmt.Printf("- Titular: %s\n", conta.Titular)
	fmt.Printf("- Número da Agência: %d\n", conta.NumeroAgencia)
	fmt.Printf("- Número da Conta: %d\n", conta.NumeroConta)
	fmt.Printf("- Saldo: %.2f\n", conta.GetSaldo())
}

func printContaPoupanca(conta contas.ContaPoupanca) {
	fmt.Println("\nDados da Conta Bancária Poupanca:")
	fmt.Printf("- Titular: %s\n", conta.Titular)
	fmt.Printf("- Número da Agência: %d\n", conta.NumeroAgencia)
	fmt.Printf("- Número da Conta: %d\n", conta.NumeroConta)
	fmt.Printf("- Saldo: %.2f\n", conta.GetSaldo())
}

func pagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) bool
}
