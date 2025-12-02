package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"loja/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	products := models.GetProducts()
	templates.ExecuteTemplate(w, "Index", products)

}

func New(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			println("Erro na conversão do preço:", err.Error())
			panic(err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			println("Erro na conversão da quantidade:", err.Error())
			panic(err.Error())
		}

		models.CreateProduct(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")
	models.DeleteProduct(idProduto)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func Edit(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")
	produto := models.EditProduct(idProduto)
	templates.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			println("Erro na conversão do ID:", err.Error())
			panic(err.Error())
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			println("Erro na conversão do preço:", err.Error())
			panic(err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			println("Erro na conversão da quantidade:", err.Error())
			panic(err.Error())
		}

		models.UpdateProduct(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
