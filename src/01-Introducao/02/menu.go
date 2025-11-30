package main

import (
	"fmt"
	"os"
)

func main() {
	exibirMenu()

	option := lerComando()

	switch option {
	case 1:
		fmt.Println("\nIniciando Monitoramento...")
	case 2:
		fmt.Println("\nExibindo Logs...")
	case 0:
		fmt.Println("\nSaindo do Programa...")
		os.Exit(0)
	default:
		fmt.Println("\nOpcao Invalida! Tente novamente.")
		os.Exit(-1)
	}
}

func exibirMenu() {
	fmt.Println("Bem-vindo ao Projeto! \nEscolha uma opcao abaixo para prosseguir:")

	fmt.Println("\n1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func lerComando() int {
	var option int
	fmt.Scan(&option)

	return option
}
