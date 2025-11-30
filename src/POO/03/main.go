package main

import (
	"POO/03/contas"
	"fmt"
)

func main() {

	conta := contas.ContaCorrente{
		Titular:       "João Silva",
		NumeroAgencia: 1234,
		NumeroConta:   567890,
		Saldo:         100.75,
	}

	fmt.Println("Dados da Conta Bancária:")
	fmt.Printf("- Titular: %s\n", conta.Titular)
	fmt.Printf("- Número da Agência: %d\n", conta.NumeroAgencia)
	fmt.Printf("- Número da Conta: %d\n", conta.NumeroConta)
	fmt.Printf("- Saldo: %.2f\n", conta.Saldo)

	sucesso := conta.Sacar(50.00)
	if sucesso {
		fmt.Printf("Saque realizado com sucesso! Novo saldo: %.2f\n", conta.Saldo)
	} else {
		fmt.Println("Saldo insuficiente para o saque.")
	}

	sucesso = conta.Depositar(150.00)
	if sucesso {
		fmt.Printf("Depósito realizado com sucesso! Novo saldo: %.2f\n", conta.Saldo)
	} else {
		fmt.Println("Valor de depósito inválido.")
	}

	conta2 := contas.ContaCorrente{
		Titular:       "Maria Oliveira",
		NumeroAgencia: 4321,
		NumeroConta:   987654,
		Saldo:         200.00,
	}

	sucesso = conta.Transferencia(75.00, &conta2)
	if sucesso {
		fmt.Printf("Transferência realizada com sucesso! Novo saldo da conta de %s: %.2f\n", conta.Titular, conta.Saldo)
		fmt.Printf("Novo saldo da conta de %s: %.2f\n", conta2.Titular, conta2.Saldo)
	} else {
		fmt.Println("Saldo insuficiente para a transferência.")
	}
}
