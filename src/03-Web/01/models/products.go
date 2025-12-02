package models

import "loja/db"

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetProducts() []Product {
	database := db.ConnectDatabase()

	rows, err := database.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, p)
	}

	defer database.Close()
	return products
}

func CreateProduct(nome string, descricao string, preco float64, quantidade int) {
	database := db.ConnectDatabase()

	insertDados, err := database.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDados.Exec(nome, descricao, preco, quantidade)
	defer database.Close()
}

func DeleteProduct(id string) {
	database := db.ConnectDatabase()

	deleteProduto, err := database.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduto.Exec(id)
	defer database.Close()
}

func EditProduct(id string) Product {
	database := db.ConnectDatabase()

	produtoDoBanco, err := database.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Product{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	defer database.Close()
	return produtoParaAtualizar
}

func UpdateProduct(id int, nome string, descricao string, preco float64, quantidade int) {
	database := db.ConnectDatabase()

	updateProduto, err := database.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduto.Exec(nome, descricao, preco, quantidade, id)
	defer database.Close()
}
