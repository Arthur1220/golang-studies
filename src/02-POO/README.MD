# 🏗️ Orientação a Objetos e Estruturas

Como Go implementa conceitos de POO utilizando Structs, Interfaces e Composição.

---

## Structs

Go não tem classes. Go tem **Structs**.
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