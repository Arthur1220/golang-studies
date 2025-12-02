package controllers

import (
	"html/template"
	"log"
	"loja/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Pagina Index acessada!")
	products := models.GetProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessou formulário de Novo Produto")
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		log.Printf("Recebida requisição de cadastro: %s", nome)

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro conversão preço:", err)
			http.Error(w, "Preço inválido", http.StatusBadRequest)
			return
		}

		qtdConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro conversão quantidade:", err)
			http.Error(w, "Quantidade inválida", http.StatusBadRequest)
			return
		}

		models.CreateProduct(nome, descricao, precoConv, qtdConv)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	log.Printf("Recebida requisição de Delete para ID: %s", idProduto)

	models.DeleteProduct(idProduto)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	log.Printf("Carregando edição para ID: %s", idProduto)

	produto := models.EditProduct(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		log.Printf("Recebida requisição de Update para ID: %s", id)

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID:", err)
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
			http.Error(w, "Preço inválido", http.StatusBadRequest)
			return
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
			http.Error(w, "Quantidade inválida", http.StatusBadRequest)
			return
		}

		models.UpdateProduct(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
