package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	for {
		exibirMenu()

		option := lerComando()

		switch option {
		case 1:
			fmt.Println("\nIniciando Monitoramento...")
			iniciarMonitoramento()
			fmt.Println("\nFim do Monitoramento...")
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

func iniciarMonitoramento() {
	var sites = []string{
		"https://brcosimple.netlify.app/",
		"https://portfolioama.netlify.app/",
		"https://www.google.com/",
		"https://www.github.com/",
	}

	for i := 0; i < monitoramentos; i++ {
		fmt.Printf("\nMonitoramento %d\n", i+1)
		for _, site := range sites {
			testarSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func testarSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Printf("\nSite: %s foi carregado com sucesso!", site)
		fmt.Printf("\nStatus Code: %d\n", response.StatusCode)
	} else {
		fmt.Printf("\nSite: %s está com problemas.", site)
		fmt.Printf("\nStatus Code: %d\n", response.StatusCode)
	}
}
