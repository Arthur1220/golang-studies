package main

import "fmt"

func main() {
	fmt.Println("Salve Familia!")

	var nome string = "Arthur"
	fmt.Println("\nOla, ", nome, "!")
	fmt.Print("Ola, " + nome + "!\n")
	fmt.Printf("Ola, %s!\n", nome)

	var versao float32 = 1.1
	fmt.Println("\nVersao do sistema:", versao)
	fmt.Printf("Versao do sistema: %v\n", versao)
	fmt.Printf("Versao do sistema: %.2f\n", versao)
	fmt.Printf("Versao do sistema: %f\n", versao)

	var idade int = 23
	fmt.Printf("\nIdade: %d anos\n", idade)

	fmt.Printf("\nOla, %s! Voce tem %d anos e esta usando a versao %.2f do sistema.\n", nome, idade, versao)
}
