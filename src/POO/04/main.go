package main

import (
	"POO/04/clientes"
	"POO/04/contas"
	"fmt"
)

func main() {

	conta := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "João Silva",
			CPF:       "123.456.789-00",
			Profissao: "Engenheiro",
		},
		NumeroAgencia: 1234,
		NumeroConta:   567890,
	}

	fmt.Println("Dados da Conta Bancária:")
	fmt.Printf("- Titular: %s\n", conta.Titular)
	fmt.Printf("- Número da Agência: %d\n", conta.NumeroAgencia)
	fmt.Printf("- Número da Conta: %d\n", conta.NumeroConta)
	fmt.Printf("- Saldo: %.2f\n", conta.GetSaldo())

	sucesso := conta.Sacar(50.00)
	if sucesso {
		fmt.Printf("Saque realizado com sucesso! Novo saldo: %.2f\n", conta.GetSaldo())
	} else {
		fmt.Println("Saldo insuficiente para o saque.")
	}

	sucesso = conta.Depositar(150.00)
	if sucesso {
		fmt.Printf("Depósito realizado com sucesso! Novo saldo: %.2f\n", conta.GetSaldo())
	} else {
		fmt.Println("Valor de depósito inválido.")
	}

	conta2 := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Maria Oliveira",
			CPF:       "987.654.321-00",
			Profissao: "Advogada",
		},
		NumeroAgencia: 4321,
		NumeroConta:   987654,
	}

	sucesso = conta.Transferencia(75.00, &conta2)
	if sucesso {
		fmt.Printf("Transferência realizada com sucesso! Novo saldo da conta de %s: %.2f\n", conta.Titular, conta.GetSaldo())
		fmt.Printf("Novo saldo da conta de %s: %.2f\n", conta2.Titular, conta2.GetSaldo())
	} else {
		fmt.Println("Saldo insuficiente para a transferência.")
	}
}
