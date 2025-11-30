package main

import "fmt"

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	conta := contaCorrente{
		titular:       "João Silva",
		numeroAgencia: 1234,
		numeroConta:   567890,
		saldo:         100.75,
	}

	fmt.Println("Dados da Conta Bancária:")
	fmt.Printf("- Titular: %s\n", conta.titular)
	fmt.Printf("- Número da Agência: %d\n", conta.numeroAgencia)
	fmt.Printf("- Número da Conta: %d\n", conta.numeroConta)
	fmt.Printf("- Saldo: %.2f\n", conta.saldo)

	sucesso := sacar(&conta, 200.00)
	if sucesso {
		fmt.Printf("Saque realizado com sucesso! Novo saldo: %.2f\n", conta.saldo)
	} else {
		fmt.Println("Saldo insuficiente para o saque.")
	}

	sucesso2 := conta.sacar2(50.00)
	if sucesso2 {
		fmt.Printf("Saque realizado com sucesso! Novo saldo: %.2f\n", conta.saldo)
	} else {
		fmt.Println("Saldo insuficiente para o saque.")
	}
}

func sacar(conta *contaCorrente, valor float64) bool {
	if valor > 0 && conta.saldo >= valor {
		conta.saldo -= valor
		return true
	}
	return false
}

func (conta *contaCorrente) sacar2(valor float64) bool {
	if valor > 0 && conta.saldo >= valor {
		conta.saldo -= valor
		return true
	}
	return false
}
