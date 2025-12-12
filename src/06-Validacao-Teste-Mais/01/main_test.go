package main

import (
	"ValidacaoTesteMais/controllers"
	"ValidacaoTesteMais/database"
	"ValidacaoTesteMais/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var ID int64

func SetupRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	// gin.DefaultWriter = ioutil.Discard
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int64(aluno.ID)
	println("ID do aluno criado no mock:", ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeSaudacao(t *testing.T) {
	rotas := SetupRotasTeste()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	rotas.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Guilherme", nil)
	w := httptest.NewRecorder()
	rotas.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	mockResposta := `{"API diz:":"E ai Guilherme, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResposta, string(respostaBody))
}

func TestListandoTodosAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	rotas := SetupRotasTeste()
	rotas.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	w := httptest.NewRecorder()
	rotas.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	rotas := SetupRotasTeste()
	rotas.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	w := httptest.NewRecorder()
	rotas.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	rotas := SetupRotasTeste()
	rotas.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	path := "/alunos/" + strconv.Itoa(int(ID))
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	rotas.ServeHTTP(w, req)
	var alunoMock models.Aluno
	json.Unmarshal(w.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	rotas := SetupRotasTeste()
	rotas.DELETE("/alunos/:id", controllers.DeletaAluno)
	path := "/alunos/" + strconv.Itoa(int(ID))
	req, _ := http.NewRequest("DELETE", path, nil)
	w := httptest.NewRecorder()
	rotas.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	rotas := SetupRotasTeste()
	rotas.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome Editado Teste", CPF: "10987654321", RG: "987654321"}
	alunoJson, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(int(ID))
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(alunoJson))
	w := httptest.NewRecorder()
	rotas.ServeHTTP(w, req)
	var alunoEditado models.Aluno
	json.Unmarshal(w.Body.Bytes(), &alunoEditado)
	assert.Equal(t, "Nome Editado Teste", alunoEditado.Nome)
	assert.Equal(t, "10987654321", alunoEditado.CPF)
	assert.Equal(t, "987654321", alunoEditado.RG)
	assert.Equal(t, http.StatusOK, w.Code)
}
