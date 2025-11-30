/*
### 🔴 Nível 4: Desafio "Real World" (Parser de Log)

**Foco:** File I/O (os/io), String Manipulation, Tratamento de Erro.

**Cenário:**
Você tem um arquivo `server.log` (pode criar esse arquivo manualmente na pasta) com o seguinte conteúdo misturado:

```text
INFO: Servidor iniciado na porta 8080
ERROR: Falha na conexão com o banco de dados
INFO: Usuário logado: admin
ERROR: Timeout na requisição API externa
INFO: Health check OK
```

**Tarefa:**
Escreva um programa que:

1.  Leia esse arquivo `server.log`.
2.  Filtre apenas as linhas que contêm a palavra **"ERROR"**.
3.  Escreva essas linhas filtradas em um novo arquivo chamado `errors.txt`.

**Dica de Go:**

  * Use `os.ReadFile` para ler tudo ou `bufio.NewScanner` para ler linha a linha (mais eficiente e profissional).
  * Use `strings.Contains(linha, "ERROR")` para verificar.
  * Use `os.OpenFile` com flags `os.O_APPEND|os.O_CREATE|os.O_WRONLY` para criar/escrever no arquivo novo.

-----
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Desafio do Log Parser!")
	linhas := lerArquivo()
	filtrarErros(linhas)
	fmt.Println("Linhas de erro foram escritas em errors.txt")

}

func lerArquivo() []string {
	var linhas []string

	arquivo, err := os.Open("server.log")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewScanner(arquivo)
	for leitor.Scan() {
		linha := leitor.Text()
		linhas = append(linhas, linha)
	}

	arquivo.Close()

	return linhas
}

func filtrarErros(linhas []string) {
	var erros []string

	for _, linha := range linhas {
		if strings.Contains(linha, "ERROR") {
			erros = append(erros, linha)
		}
	}

	escreverErros(strings.Join(erros, "\n"))
}

func escreverErros(linha string) {

	arquivo, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(linha + "\n")

	arquivo.Close()
}
