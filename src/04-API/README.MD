# 🏛️ API de Personalidades Históricas

API RESTful desenvolvida em Go para gerenciar um catálogo de personalidades históricas. O projeto demonstra a construção de um CRUD completo conectado a um banco de dados PostgreSQL, utilizando Docker para orquestração de ambiente.

## 🚀 Tecnologias e Bibliotecas

  * **[Gorilla Mux](https://github.com/gorilla/mux):** Roteador HTTP robusto e flexível.
  * **[GORM](https://gorm.io/):** ORM (Object Relational Mapper) para facilitar a interação com o banco de dados.
  * **[PostgreSQL](https://www.postgresql.org/):** Banco de dados relacional.
  * **Docker & Docker Compose:** Para subir o banco e o pgAdmin automaticamente.

## ⚙️ Como Executar o Projeto

### 1\. Pré-requisitos

Certifique-se de ter o [Docker](https://www.docker.com/) e o [Go](https://go.dev/) instalados.

### 2\. Subindo o Banco de Dados

O projeto utiliza um arquivo `docker-compose.yml` que sobe o PostgreSQL e já cria a tabela e insere dados iniciais (via script de migração).

```bash
# Na raiz do projeto, execute:
docker-compose up -d
```

*Isso iniciará o Postgres na porta `5432` e o pgAdmin na porta `54321`.*

### 3\. Rodando a API

```bash
# Baixe as dependências
go mod tidy

# Execute o servidor
go run main.go
```

O servidor iniciará em: **http://localhost:8000**

-----

## 🔗 Endpoints da API

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| `GET` | `/api/personalidades` | Retorna todas as personalidades (JSON) |
| `GET` | `/api/personalidades/{id}` | Retorna uma personalidade específica pelo ID |
| `POST` | `/api/personalidades` | Cria uma nova personalidade |
| `PUT` | `/api/personalidades/{id}` | Edita/Atualiza uma personalidade existente |
| `DELETE` | `/api/personalidades/{id}` | Remove uma personalidade do banco |

-----

## 🧠 Análise Técnica e Destaques do Código

Abaixo, detalho as decisões arquiteturais e trechos importantes do código.

### 1\. ORM e Abstração do Banco (`database/db.go`)

Utilizamos o **GORM** para não precisarmos escrever SQL puro manualmente. Note como a conexão é simples:

```go
// A string de conexão define o host, usuário, senha e banco (configurados no Docker)
stringDB := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
DB, err = gorm.Open(postgres.Open(stringDB))
```

### 2\. Controllers e JSON (`controllers/controllers.go`)

O Go trata JSON nativamente. Nos controllers, fazemos o processo de **Encoding** (Go -\> JSON) e **Decoding** (JSON -\> Go).

**Exemplo: Criando um registro (POST)**

```go
func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
    var personalidade models.Personalidade
    
    // 1. Decoder: Pega o JSON que veio no corpo da requisição (Body) e converte para a Struct
    json.NewDecoder(r.Body).Decode(&personalidade)
    
    // 2. GORM: O método .Create insere no banco automaticamente
    database.DB.Create(&personalidade)
    
    // 3. Encoder: Devolve o objeto criado como JSON para o cliente
    json.NewEncoder(w).Encode(personalidade)
}
```

**Exemplo: Buscando por ID (GET)**
O `Gorilla Mux` facilita a extração de variáveis da URL:

```go
vars := mux.Vars(r) // Pega todas as variáveis da rota
id := vars["id"]    // Extrai o ID da URL /api/personalidades/{id}

database.DB.First(&personalidade, id) // O GORM busca automaticamente pelo Primary Key
```

### 3\. Middleware (`middleware/middleware.go`)

Middlewares são funções que interceptam a requisição antes dela chegar no Controller. Criamos um middleware para garantir que **todas** as respostas da API tenham o cabeçalho `Content-Type: application/json`.

```go
func ContentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json") // Define JSON
        next.ServeHTTP(w, r) // Passa para o próximo passo (o Controller)
    })
}
```

### 4\. Roteamento (`routes/routes.go`)

Aqui amarramos tudo. Definimos as rotas, os verbos HTTP permitidos e injetamos o Middleware.

```go
r := mux.NewRouter()
r.Use(middleware.ContentTypeMiddleware) // Aplica o middleware globalmente

// Definimos explicitamente qual verbo HTTP cada função aceita
r.HandleFunc("/api/personalidades", controllers.CriarPersonalidade).Methods("POST")
```

-----

## 🐳 Docker e Migração Automática

O arquivo `migration/docker-database-initial.sql` é montado dentro do container do Postgres. Quando o container nasce pela primeira vez, o Postgres executa esse script automaticamente:

```sql
create table personalidades( ... );
INSERT INTO personalidades ...;
```

Isso garante que qualquer pessoa que baixe seu projeto tenha o banco pronto para uso sem precisar criar tabelas manualmente.

-----

Desenvolvido com Go 🐹