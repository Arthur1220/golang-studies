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
		saldo:         1500.75,
	}

	fmt.Println("Dados da Conta Bancária:")
	fmt.Printf("- Titular: %s\n", conta.titular)
	fmt.Printf("- Número da Agência: %d\n", conta.numeroAgencia)
	fmt.Printf("- Número da Conta: %d\n", conta.numeroConta)
	fmt.Printf("- Saldo: %.2f\n", conta.saldo)
}
