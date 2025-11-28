# 📘 Guia de Estudos e Setup: Go (Golang)

![Go Version](https://img.shields.io/badge/Go-1.25.4-blue) ![Status](https://img.shields.io/badge/Status-Learning-green)

Espaço destinado aos meus estudos e experimentos com **Go**.
Este repositório serve como um guia pessoal de consulta e portfólio de evolução na linguagem.

### 📚 Recursos Oficiais
Para quem também está estudando, recomendo:
- [Documentação Oficial](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go Playground](https://go.dev/play/)


-----

## 🛠️ Instalação e Configuração (WSL/Linux)

Passo a passo para instalar a versão mais recente do Go no Linux (Ubuntu/WSL).

### 1\. Limpeza e Download

```bash
# Remove instalações anteriores para evitar conflitos
sudo rm -rf /usr/local/go

# Baixa a versão estável (Verifique a versão atual em go.dev/dl)
wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz

# Extrai para o diretório padrão
sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
```

### 2\. Configuração do PATH

Dependendo do seu terminal (`bash` ou `zsh`), execute os comandos abaixo para que o sistema encontre o comando `go`.

**Para Bash:**

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

**Para Zsh (Padrão em muitos setups modernos):**

```zsh
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

### 3\. Verificação e Inicialização

```bash
# Verifica se instalou corretamente
go version

# Inicializa o gerenciador de dependências (cria o arquivo go.mod)
go mod init nome-do-projeto
```

-----

## 📂 Estrutura e Compilação

### Organização de Pastas (Convenção Clássica)

Embora o Go Modules tenha flexibilizado isso, a estrutura clássica é:

  * `bin/`: Binários compilados (executáveis).
  * `pkg/`: Objetos de pacotes compilados (bibliotecas).
  * `src/`: Código fonte (`.go`).

### O Arquivo Principal

Todo programa executável em Go deve começar definindo o pacote como `main`. É similar à função `main` do C: o ponto de entrada.

```go
package main

import "fmt" // Pacote para I/O (Input/Output) formatado

func main() {
    // Código aqui
}
```

### Rodando o Código

Go é uma linguagem compilada (como C), mas pode ser rodada diretamente para testes.

  * **Compilar e Executar (Rápido):** `go run arquivo.go`
  * **Apenas Compilar (Gera executável):** `go build arquivo.go`
  * **Executar binário:** `./arquivo`

-----

## 📝 Sintaxe Básica e Variáveis

### Declaração de Variáveis

Go é fortemente tipado, mas possui inferência de tipo (parecido com Python, mas estático).

```go
// 1. Declaração explícita (Estilo C)
var nome string = "João"

// 2. Inferência de tipo (O compilador descobre)
var idade = 30

// 3. Short Syntax (A mais usada dentro de funções)
email := "joao@email.com" 
```

> **Regra Importante:** O Go não compila se houver variáveis declaradas e não utilizadas (evita lixo no código).
> **Valor Zero:** Se não atribuir valor, Go define um padrão: `0` para números, `""` (vazio) para strings, `false` para booleanos.

### Descobrindo o Tipo

```go
import "reflect"
// ...
fmt.Println(reflect.TypeOf(nome_da_variavel))
```

### Tipos Primitivos Comuns

  * `string`
  * `int`, `int8`, `int64`
  * `float32`, `float64`
  * `bool`

### Constantes

Valores imutáveis, definidos em tempo de compilação.

```go
const Pi = 3.14
const Status = "Ativo"
```

-----

## 🖨️ Output e Input (Prints e Scans)

O pacote `fmt` é o padrão.

### Saída (Output)

  * `fmt.Print`: Imprime sem pular linha.
  * `fmt.Println`: Imprime e pula linha no final.
  * `fmt.Printf`: Imprime formatado (igual ao C).

**Formatadores do Printf:**

  * `%s`: String
  * `%d`: Inteiro (Decimal)
  * `%f`: Float (`%.2f` para 2 casas decimais)
  * `%v`: Valor genérico (serve para qualquer coisa, muito útil)
  * `%T`: Mostra o Tipo da variável

### Entrada (Input)

```go
var comando int

// Opção 1: Scanf (Estilo C - precisa formatador)
fmt.Scanf("%d", &comando) 

// Opção 2: Scan (Mais inteligente - infere o tipo)
fmt.Scan(&comando)
```

> O `&` é o ponteiro, indicando o endereço de memória onde o valor será salvo.

-----

## 🔀 Controle de Fluxo

### IF / ELSE

Não usa parênteses na condição, mas as chaves `{}` são obrigatórias.

```go
if idade >= 18 {
    fmt.Println("Maior de idade")
} else {
    fmt.Println("Menor de idade")
}
```

### Switch

Não precisa de `break` (ele sai automaticamente após executar o case, diferente de C que executaria os próximos se não parasse).

```go
switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 0:
        fmt.Println("Saindo...")
    default:
        fmt.Println("Comando desconhecido")
}
```

### Loops (For)

Go **não** tem `while`. O `for` faz tudo.

```go
// 1. Loop Infinito (Substituto do while(true))
for {
    // Roda para sempre até um break
}

// 2. Loop Clássico (Estilo C)
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// 3. Range (Estilo Python enumerate)
nomes := []string{"Ana", "Jose"}
for i, nome := range nomes {
    fmt.Println(i, nome)
}
// Se não precisar do índice, use o Blank Identifier (_)
for _, nome := range nomes {
    fmt.Println(nome)
}
```

-----

## 📦 Arrays e Slices (Listas)

### Arrays

Tamanho fixo. Pouco usado na prática, pois é rígido (igual C).

```go
var sites [4]string
sites[0] = "google.com"
```

### Slices

Tamanho dinâmico (abstração em cima do array). Muito usado (parecido com listas Python).

```go
// Declaração
nomes := []string{"Douglas", "Daniel"}

// Adicionando elementos (Função append)
nomes = append(nomes, "Novo Nome")

// Propriedades
len(nomes) // Tamanho atual (Length)
cap(nomes) // Capacidade do array subjacente
```

> Quando a capacidade estoura, o Go automaticamente dobra o tamanho do array interno.

-----

## 🧩 Funções

Sintaxe: `func nome(parametros) retorno`.
Go permite **múltiplos retornos**, uma característica poderosa.

```go
// Função que retorna duas informações
func devolverNomeEIdade() (string, int) {
    return "Arthur", 25
}

func main() {
    // Recebendo os dois valores
    nome, idade := devolverNomeEIdade()
    
    // Ignorando um valor com _
    _, idadeSozinha := devolverNomeEIdade()
}
```

-----

## ⚙️ Funcionalidades de Sistema e Web

### Sair do Programa

```go
import "os"

os.Exit(0)  // Sucesso
os.Exit(-1) // Erro
```

### Limpar o Terminal

Go não tem comando nativo direto como `system("cls")` simples, depende do SO. A forma correta usando o pacote `os/exec`:

```go
import ("os"; "os/exec")

func limparTerminal() {
    c := exec.Command("clear") // Use "cls" se estiver rodando no CMD do Windows nativo
    c.Stdout = os.Stdout
    c.Run()
}
```

### Requisições Web (HTTP)

```go
import "net/http"

func testaSite(site string) {
    // O _ ignora o erro (não recomendado em prod, mas ok para teste simples)
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site carregado!")
    }
}
```

### O que é `defer`?

O `defer` adia a execução de uma função até o momento que a função atual terminar (return). É perfeito para limpeza de recursos.

```go
func lerSite() {
    resp, _ := http.Get("http://google.com")
    
    // Assim que a função lerSite acabar, feche a conexão.
    // Isso garante que não vamos esquecer a conexão aberta, mesmo se der erro no meio.
    defer resp.Body.Close() 
    
    // ... código de leitura ...
}
```

### Time / Sleep

```go
import "time"

// Pausa a execução por 5 segundos
time.Sleep(5 * time.Second)
```

-----

## 📄 Arquivos (Files)

Para ler e escrever arquivos, usamos os pacotes `os` (operações de sistema) e `io` ou `bufio` (leitura).

### Lendo um arquivo inteiro

```go
import ("os"; "fmt")

func lerArquivo() {
    // Retorna um array de bytes (conteúdo) e um erro
    arquivo, err := os.ReadFile("log.txt") 
    
    if err != nil {
        fmt.Println("Erro ao ler:", err)
    }
    
    // Converte bytes para string
    fmt.Println(string(arquivo)) 
}
```

### Escrevendo em um arquivo

```go
func escreverLog() {
    // Abre o arquivo. Flags: Criar se não existir, Escrita, Append (adicionar no final)
    arquivo, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    
    defer arquivo.Close() // Garante o fechamento do arquivo
    
    arquivo.WriteString("Log de erro registrado\n")
}
```