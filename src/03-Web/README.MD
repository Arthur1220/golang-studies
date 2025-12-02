# 🛒 Loja Web em Go (MVC)

Aplicação Web desenvolvida durante o curso de Go da Alura, focada no desenvolvimento web, manipulação de banco de dados e arquitetura MVC. O projeto consiste em um sistema de gerenciamento de produtos (CRUD) com interface gráfica renderizada pelo servidor.

## 🚀 Funcionalidades

O sistema permite realizar as 4 operações fundamentais em um banco de dados:

  * **Listar Produtos:** Exibe uma tabela com todos os produtos cadastrados.
  * **Adicionar Produto:** Formulário para inserção de novos itens.
  * **Editar Produto:** Carrega os dados de um produto existente para atualização.
  * **Deletar Produto:** Remove um produto do banco (com confirmação via JavaScript).

## 🛠️ Tecnologias Utilizadas

  * **Linguagem:** Go (Golang)
  * **Banco de Dados:** PostgreSQL
  * **Driver:** `github.com/lib/pq`
  * **Front-end:** HTML5, CSS3 (Bootstrap 4), JavaScript (Simples)
  * **Template Engine:** `html/template` (Nativo do Go)

## 🏗️ Arquitetura do Projeto

O projeto segue o padrão **MVC (Model-View-Controller)** para separar as responsabilidades:

[Image of MVC architecture diagram web application]

  * **`/controllers`**: Gerencia as requisições HTTP, chama os modelos e define qual template exibir.
  * **`/models`**: Contém a lógica de acesso ao banco de dados e as structs dos dados.
  * **`/templates` (Views)**: Arquivos HTML que são renderizados com os dados dinâmicos.
  * **`/db`**: Responsável pela conexão com o banco de dados.
  * **`/routes`**: Define as rotas da aplicação (endpoints).

-----

## 🧠 Aprofundamento Técnico

Pontos importantes sobre a implementação e decisões de código:

### 1\. Conexão Segura com Banco (Prepared Statements)

Para evitar ataques de **SQL Injection**, não concatenamos strings diretamente nas queries. Utilizamos `Prepared Statements` (o `$1, $2` nas queries).

```go
// ❌ Ruim (Vulnerável):
// db.Query("insert into produtos values (" + nome + ")")

// ✅ Bom (Seguro):
stmt, err := db.Prepare("insert into produtos(nome) values($1)")
stmt.Exec(nome)
```

Isso garante que o driver do Postgres trate os dados antes de executar o comando.

### 2\. Gerenciamento de Recursos (Defer)

Em Go, é crucial fechar conexões de banco de dados para não vazar memória ou estourar o limite de conexões do servidor. Utilizamos a keyword `defer` para garantir que a conexão feche assim que a função terminar, independentemente se houve erro ou sucesso.

```go
database := db.ConnectDatabase()
defer database.Close() // Executa por último, garantindo a limpeza
```

### 3\. Server-Side Rendering (SSR)

Diferente de Single Page Applications (React/Vue) que recebem JSON, aqui o Go processa o HTML no servidor e devolve a página pronta.

  * **`template.ParseGlob`**: Carrega todos os arquivos `.html` na memória na inicialização (performance).
  * **`{{range .}}`**: Função do template engine do Go para iterar sobre o Slice de produtos diretamente no HTML.

### 4\. Conversão de Tipos

Como o HTML envia tudo como `string`, a camada de **Controller** atua como um tradutor, convertendo os inputs para `float64` (preço) e `int` (quantidade) antes de passar para o **Model**. Isso garante a integridade dos dados antes de chegarem ao banco.

## ⚙️ Configuração do Ambiente

### 1\. Pré-requisitos

Certifique-se de ter instalado:

  * [Go](https://go.dev/dl/) (Versão 1.18+)
  * [PostgreSQL](https://www.postgresql.org/download/)

### 2\. Configuração do Banco de Dados

O projeto está configurado para conectar em um banco local. Execute os comandos SQL abaixo no seu cliente Postgres (PgAdmin, DBeaver ou Terminal) para criar o banco e a tabela necessária:

```sql
-- 1. Criar o banco de dados
CREATE DATABASE alura_loja;

-- 2. Conectar ao banco
\c alura_loja;

-- 3. Criar a tabela de produtos
CREATE TABLE produtos (
    id serial primary key,
    nome varchar(255),
    descricao varchar(255),
    preco decimal,
    quantidade integer
);

-- 4. (Opcional) Inserir dados iniciais
INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES
('Camiseta', 'Camiseta de algodão', 39.90, 10),
('Notebook', 'Notebook i7', 4500.00, 5);
```

> **Nota:** As credenciais de acesso ao banco estão definidas em `db/db.go`. Caso sua senha ou usuário sejam diferentes de `postgres/123`, altere a string de conexão neste arquivo.

### 3\. Instalação e Execução

Clone o repositório e navegue até a pasta do projeto:

```bash
# Baixar as dependências (Driver do Postgres)
go mod tidy

# Rodar o servidor
go run main.go
```

Após iniciar, acesse no seu navegador:
👉 **http://localhost:8000**

## 🛣️ Rotas da Aplicação

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| `GET` | `/` | Página inicial (Lista de produtos) |
| `GET` | `/new` | Formulário para criar novo produto |
| `POST` | `/insert` | Ação que salva o novo produto no banco |
| `GET` | `/delete?id=X` | Remove o produto de ID X |
| `GET` | `/edit?id=X` | Formulário de edição preenchido |
| `POST` | `/update` | Ação que atualiza o produto no banco |

## 📦 Estrutura de Pastas

```bash
03-Web/
├── controllers/   # Lógica de controle (handlers)
├── db/            # Conexão com PostgreSQL
├── models/        # Structs e Queries SQL
├── routes/        # Definição de endpoints
├── templates/     # Arquivos HTML (Index, New, Edit)
├── main.go        # Ponto de entrada do servidor
├── go.mod         # Gerenciador de dependências
└── README.MD      # Documentação
```

-----

Desenvolvido como parte dos estudos em Go. 🐹