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

### Zero-Initialization (Valores Padrão)

Em Go, variáveis não inicializadas **não** contêm lixo de memória. Elas começam com valores "Zero":

| Tipo | Valor Zero |
| :--- | :--- |
| `int`, `float` | `0` |
| `string` | `""` (vazio) |
| `bool` | `false` |
| `pointer`, `interface`, `slice` | `nil` |
| `struct` | `{}` (campos zerados) |

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

## 🖨️ Output e Input

### Saída (Output)

  * `fmt.Print`: Imprime sem pular linha.
  * `fmt.Println`: Imprime e pula linha no final.
  * `fmt.Printf`: Imprime formatado.

**Formatadores do Printf:**

  * `%s`: String
  * `%d`: Inteiro (Decimal)
  * `%f`: Float (`%.2f` para 2 casas decimais)
  * `%v`: Valor genérico (serve para qualquer coisa)
  * `%T`: Mostra o Tipo da variável

### Entrada (Input)

```go
var comando int
// O & indica o endereço de memória onde o valor será salvo.
fmt.Scan(&comando)
```

-----

## 🔀 Controle de Fluxo

### IF / ELSE

Sem parênteses na condição, chaves `{}` obrigatórias.

```go
if idade >= 18 {
    fmt.Println("Maior de idade")
} else {
    fmt.Println("Menor de idade")
}
```

### Switch

Não precisa de `break` (automático).

```go
switch comando {
    case 1:
        fmt.Println("Monitorando...")
    default:
        fmt.Println("Comando desconhecido")
}
```

### Loops (For)

Go **não** tem `while`. O `for` faz tudo.

```go
// 1. Loop Infinito (While true)
for { }

// 2. Loop Clássico (C Style)
for i := 0; i < 10; i++ { }

// 3. Range (Python Style)
nomes := []string{"Ana", "Jose"}
for i, nome := range nomes {
    fmt.Println(i, nome)
}
```

-----

## 📦 Arrays e Slices

### Arrays

Tamanho fixo. Pouco usado.

```go
var sites [4]string
```

### Slices

Tamanho dinâmico. Muito usado.

```go
nomes := []string{"Douglas", "Daniel"}
nomes = append(nomes, "Novo Nome")

len(nomes) // Tamanho atual
cap(nomes) // Capacidade total alocada
```

> Quando a capacidade estoura, o Go dobra o tamanho do array interno automaticamente.

-----

## ⚙️ Funcionalidades de Sistema e Web

### Time / Sleep

Podemos pausar a execução da *goroutine* atual utilizando a função `Sleep` do pacote `time`.
Isso é essencial em loops de monitoramento para não sobrecarregar a CPU ou o serviço que está sendo consultado, criando um intervalo saudável entre as ações.

```go
import "time"

// Pausa a execução por 5 segundos
// time.Second é uma constante que representa 1 segundo
time.Sleep(5 * time.Second)
```

### Requisições Web (HTTP)

Para fazer chamadas a APIs ou sites externos, usamos o pacote nativo `net/http`.

```go
import ("net/http"; "fmt")

func testaSite(site string) {
    // O _ ignora o erro (não recomendado em prod, mas ok para teste simples)
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site carregado com sucesso!")
    }
}
```

### Sair e Limpar

```go
import ("os"; "os/exec")

// Sai do programa com código de sucesso (0) ou erro (-1)
os.Exit(0) 

// Limpar Terminal (Linux/Mac)
func limpar() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}
```

### Defer

Adia a execução de uma função até o momento que a função pai terminar (return). É perfeito para garantir o fechamento de recursos.

```go
func ler() {
    arquivo, _ := os.Open("dados.txt")
    defer arquivo.Close() // Será executado por último, não importa o que aconteça
    // ... processamento ...
}
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
    // 0666 é a permissão do arquivo no Linux (Leitura/Escrita)
    arquivo, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    
    defer arquivo.Close() // Garante o fechamento do arquivo
    
    arquivo.WriteString("Log de erro registrado\n")
}
```

-----

## 🏗️ Orientação a Objetos em Go

Go não tem classes. Go tem **Structs**.

### Structs

É uma coleção de campos. Funciona como o `struct` de C ou objetos literais de JS.

```go
type ContaCorrente struct {
    Titular       string
    NumeroAgencia int
    NumeroConta   int
    Saldo         float64
}

// Instanciando
conta := ContaCorrente{
    Titular:       "João Silva",
    NumeroAgencia: 1234,
    Saldo:         1500.75, // Campos omitidos ficam com Valor Zero
}

// Acessando
fmt.Println(conta.Titular)
```

### Ponteiros e Alocação (`new` vs `&`)

Ponteiros armazenam o **endereço de memória** de um valor, não o valor em si. Isso economiza memória (evita cópias de objetos grandes) e permite alterar o valor original.

**1. Usando `&` (Mais comum):**
Cria a estrutura e retorna o endereço dela.

```go
var conta *ContaCorrente
conta = &ContaCorrente{Titular: "Maria"} 
```

**2. Usando `new`:**
Aloca memória zerada e retorna o ponteiro. Pouco usado em estruturas complexas pois não permite inicializar campos na hora.

```go
conta2 := new(ContaCorrente) // Retorna *ContaCorrente
conta2.Titular = "José"
```

-----

## ⚡ Métodos e Funções

### Funções Variádicas

Funções que aceitam um número indeterminado de argumentos (usando `...`).

```go
func Somando(numeros ...int) int {
    soma := 0
    for _, num := range numeros {
        soma += num
    }
    return soma
}
// Uso: Somando(1, 2, 3, 4, 5)
```

### Métodos vs. Funções (Receivers)

Em Go, podemos "anexar" uma função a uma Struct. Isso é o equivalente a **Métodos** em outras linguagens.

**1. Função Tradicional (Style C):**
Recebe a struct como parâmetro. Não faz parte do objeto.

```go
func sacar(conta *ContaCorrente, valor float64) { ... }
// Uso: sacar(&conta, 200)
```

**2. Método com Receiver (Style Go/POO):**
A função "pertence" à struct. O `(c *ContaCorrente)` antes do nome da função é o **Receiver**.

```go
func (c *ContaCorrente) Sacar(valor float64) string {
    if c.Saldo >= valor {
        c.Saldo -= valor
        return "Sucesso"
    }
    return "Saldo Insuficiente"
}
// Uso: conta.Sacar(200)
```

> **Dica:** Use **Pointer Receiver** (`*ContaCorrente`) se você precisa alterar o valor da struct original. Se usar Value Receiver (`ContaCorrente`), você estará mexendo em uma cópia.

-----

## 📦 Pacotes e Visibilidade

Para organizar o código, dividimos em pastas (pacotes).

### Comandos Essenciais

  * `go mod init <nome>`: Cria o módulo inicial.
  * `go mod tidy`: Baixa dependências que faltam e remove as que não são usadas (limpeza).
  * `go build ./...`: Compila todos os pacotes dentro da pasta para verificar erros.

### Encapsulamento (Public vs Private)

Go não tem `public` ou `private`. A visibilidade é definida pela **Letra Inicial**.

  * **Maiúscula (ex: `Sacar`)**: **Público (Exported)**. Visível para outros pacotes.
  * **Minúscula (ex: `saldo`)**: **Privado (Unexported)**. Visível apenas dentro do pacote onde foi criado.

**Exemplo de Encapsulamento (Getters/Setters):**

```go
package contas

type Conta struct {
    Titular string  // Público
    saldo   float64 // Privado (ninguém mexe direto)
}

// Método Público para ler o saldo privado
func (c *Conta) ObterSaldo() float64 {
    return c.saldo
}
```

-----

## 🔁 Interfaces (Polimorfismo)

Interfaces em Go definem **comportamento**. Se uma struct tem os métodos que a interface pede, ela implementa a interface automaticamente (Duck Typing). Não existe `implements`.

```go
// Contrato: "Qualquer coisa que tenha o método VerificarConta é um Verificavel"
type Verificavel interface {
    VerificarConta() bool
}

func (c ContaCorrente) VerificarConta() bool {
    return true
}

// Esta função aceita QUALQUER coisa que respeite a interface Verificavel
func RealizarVerificacao(v Verificavel) {
    v.VerificarConta()
}
```