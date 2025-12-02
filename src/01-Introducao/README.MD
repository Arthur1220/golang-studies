# 📘 Fundamentos e Setup

Guia de instalação, configuração de ambiente e sintaxe base da linguagem Go.

---

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
