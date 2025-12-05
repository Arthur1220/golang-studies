# 🍸 API de Estudantes com Gin & GORM

Bem-vindo à evolução do desenvolvimento web em Go. Este projeto implementa uma API REST completa para gerenciamento de estudantes, utilizando o framework **Gin** (para rotas e HTTP) e **GORM** (para banco de dados).

O objetivo principal deste módulo é demonstrar o ganho de **produtividade** e **performance** ao sair do roteamento manual (`net/http` + `Gorilla Mux`) para um Framework Web moderno.

-----

## 🆚 O Poder do Gin: Comparativo

A principal diferença que você notará é a redução de código boilerplate (código repetitivo). Veja a comparação:

### 1\. Retornando JSON

**Sem Gin (net/http):**

```go
// Você precisava configurar o header e criar um Encoder manualmente
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(aluno)
```

**Com Gin:**

```go
// O Gin já sabe que é JSON, configura o header e serializa tudo em uma linha
c.JSON(200, aluno)
```

### 2\. Lendo o Corpo da Requisição (Body)

**Sem Gin:**

```go
var aluno Aluno
// Precisava criar um Decoder e tratar erros manualmente
err := json.NewDecoder(r.Body).Decode(&aluno)
if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
}
```

**Com Gin (Magic Binding):**

```go
var aluno Aluno
// ShouldBindJSON faz o decode E valida os campos automaticamente
if err := c.ShouldBindJSON(&aluno); err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
}
```

### 3\. Pegando Parâmetros da URL

**Sem Gin (Gorilla Mux):**

```go
vars := mux.Vars(r)
id := vars["id"]
```

**Com Gin:**

```go
id := c.Param("id") // Simples e direto
```

-----

## 🚀 Como Executar

### 1\. Subindo o Banco (Docker)

Este projeto usa PostgreSQL versão 16.

```bash
docker-compose up -d
```

### 2\. Rodando a API

```bash
go mod tidy
go run main.go
```

O servidor iniciará em **http://localhost:8080** (Padrão do Gin, diferente do 8000 anterior).

-----

## 🔗 Endpoints Disponíveis

| Verbo | Rota | Descrição |
| :--- | :--- | :--- |
| `GET` | `/` | Retorna mensagem de boas vindas |
| `GET` | `/introduce/:name` | Exemplo de uso de parâmetros na URL |
| `GET` | `/students` | Lista todos os estudantes |
| `POST` | `/students` | Cria um novo estudante (JSON Body) |
| `GET` | `/students/:id` | Busca estudante por ID |
| `DELETE` | `/students/:id` | Deleta estudante por ID |
| `PUT` | `/students/:id` | Atualiza dados do estudante |
| `GET` | `/students/name/:name` | Busca estudante pelo nome (Busca exata) |

-----

## 🧠 Detalhes da Implementação

### Estrutura MVC Simplificada

  * **`routes/routes.go`**: Onde definimos os endpoints. Note o uso de `gin.Default()`, que já cria um roteador com **Logger** (logs coloridos no terminal) e **Recovery** (se der panic, a API não cai, ela retorna 500).
  * **`controllers/controllers.go`**: A lógica das rotas.
  * **`models/student.go`**: Definição da tabela do banco.
  * **`database/db.go`**: Conexão com GORM.

### O Contexto do Gin (`*gin.Context`)

Diferente do Go puro que passa `w` (Writer) e `r` (Request) separados, o Gin encapsula tudo no **Contexto (`c`)**.

  * É através do `c` que você pega dados (`c.Param`, `c.Query`, `c.ShouldBindJSON`).
  * É através do `c` que você responde (`c.JSON`, `c.String`, `c.HTML`).

### GORM: Auto Migration

No arquivo `database/db.go`, utilizamos:

```go
DB.AutoMigrate(&models.Student{})
```

Isso significa que você **não precisa criar tabela SQL na mão**. O GORM olha para a sua Struct `Student` e cria (ou atualiza) a tabela no Postgres automaticamente quando o servidor sobe.

### Tratamento de Erros

O Gin facilita retornar erros amigáveis. Usamos muito o `gin.H{}` que é um atalho para `map[string]interface{}`.

```go
c.JSON(404, gin.H{
    "error": "Student not found",
})
```

-----

Desenvolvido com 💜 e Go.