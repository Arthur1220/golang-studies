# ✅ API Completa: Validação, Testes e Frontend

Este projeto representa o nível final da formação básica em Go. Aqui, expandimos uma API REST com Gin para incluir **Validação de Dados** rigorosa, **Testes Automatizados** (Unitários/Integração) e integração com **Frontend (HTML/CSS)** renderizado pelo Go.

## 🚀 Novas Competências Adquiridas

Diferente dos módulos anteriores, aqui o foco não é apenas "funcionar", mas "funcionar corretamente e com segurança".

  * **Validação de Structs:** Uso de tags para garantir que CPF e RG tenham o tamanho e formato corretos antes de salvar no banco.
  * **Testes Automatizados:** Criação de testes para garantir que as rotas retornem os status codes e dados esperados.
  * **Arquivos Estáticos:** Configuração do Gin para servir CSS e HTML (Templates).
  * **Mock de Dados:** Técnicas para testar a API sem depender 100% do banco de dados em todos os cenários.

-----

## 🛠️ Tecnologias e Bibliotecas

  * **[Gin Gonic](https://github.com/gin-gonic/gin):** Framework Web.
  * **[GORM](https://gorm.io/):** ORM para PostgreSQL.
  * **[Validator v2](https://gopkg.in/validator.v2):** Biblioteca para validação de dados via tags.
  * **[Testify](https://github.com/stretchr/testify):** (Opcional) Usado para facilitar asserções nos testes (`assert.Equal`).

-----

## 🧠 Destaques do Código

### 1\. Validação de Dados (`models/aluno.go`)

Em vez de encher o Controller de `if len(cpf) != 11`, usamos tags na Struct. O Gin e o Validator fazem o trabalho pesado.

```go
type Aluno struct {
    gorm.Model
    // Validate: não pode ser vazio (nonzero)
    Nome string `json:"nome" validate:"nonzero"`
    // Validate: deve ter tamanho 11 e apenas números regex
    CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]+$"`
    RG   string `json:"rg" validate:"len=9, regexp=^[0-9]+$"`
}

// Função auxiliar que chama o validador
func ValidateData(aluno *Aluno) error {
    if err := validator.Validate(aluno); err != nil {
        return err
    }
    return nil
}
```

### 2\. Testes Automatizados (`main_test.go`)

O Go possui um pacote nativo robusto para testes (`testing` e `net/http/httptest`).

**Exemplo de teste de rota:**

```go
func TestVerificaStatusCodeSaudacao(t *testing.T) {
    rotas := SetupRotasTeste() // Cria um router isolado
    
    // Simula uma requisição GET
    req, _ := http.NewRequest("GET", "/Guilherme", nil)
    w := httptest.NewRecorder() // Gravador de resposta
    
    rotas.ServeHTTP(w, req) // Executa a rota

    // Verifica se o código HTTP é 200 (OK)
    assert.Equal(t, http.StatusOK, w.Code)
}
```

### 3\. Servindo HTML e CSS (`routes/routes.go`)

O Gin facilita muito a configuração de arquivos estáticos comparado ao `net/http` puro.

```go
// Carrega os templates HTML da pasta
r.LoadHTMLGlob("templates/*")

// Define que a rota /assets vai servir arquivos da pasta ./assets
r.Static("/assets", "./assets")

// Rota que renderiza HTML
r.GET("/index", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
        "alunos": listaDeAlunos,
    })
})
```

-----

## ⚙️ Como Executar

### 1\. Subindo o Ambiente (Docker)

Este projeto utiliza PostgreSQL 16.

```bash
docker-compose up -d
```

### 2\. Executando a Aplicação

```bash
go mod tidy
go run main.go
```

  * **API:** http://localhost:8080/alunos
  * **Frontend:** http://localhost:8080/index (Visualização HTML)

### 3\. Executando os Testes 🧪

Para rodar a bateria de testes automatizados e verificar a integridade da API:

```bash
go test -v
```

*(O flag `-v` mostra o log detalhado de cada teste executado)*.

-----

## 📂 Estrutura do Projeto

```bash
06-Validacao-Teste-Mais/
├── assets/          # Arquivos CSS (Estilos)
├── controllers/     # Lógica das rotas
├── database/        # Conexão GORM (com TimeZone configurado)
├── models/          # Structs e Validações
├── routes/          # Configuração de rotas e arquivos estáticos
├── templates/       # Arquivos HTML (Index, 404)
├── main.go          # Entry point
├── main_test.go     # 🧪 ARQUIVO DE TESTES
└── docker-compose.yml
```

-----

Desenvolvido com 💜 e Go.