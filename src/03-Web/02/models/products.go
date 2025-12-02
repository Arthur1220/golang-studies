package models

import (
	"log"
	"loja/db"
)

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetProducts() []Product {
	db := db.ConnectDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		log.Println("Erro ao buscar produtos no banco:", err)
		return nil
	}
	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			log.Println("Erro ao scanear linha do produto:", err)
			continue
		}
		products = append(products, p)
	}
	return products
}

func CreateProduct(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectDatabase()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Println("Erro ao preparar query de inserção:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		log.Println("Erro ao executar inserção:", err)
		return
	}

	log.Printf("Novo produto cadastrado com sucesso: %s | R$ %.2f", nome, preco)
}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		log.Println("Erro ao preparar query de delete:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Erro ao deletar produto:", err)
		return
	}

	log.Printf("Produto de ID %s deletado com sucesso.", id)
}

func EditProduct(id string) Product {
	db := db.ConnectDatabase()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM produtos WHERE id=$1", id)

	var p Product
	err := row.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
	if err != nil {
		log.Println("Erro ao buscar produto para edição (ID "+id+"):", err)
		return p
	}
	return p
}

func UpdateProduct(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectDatabase()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		log.Println("Erro ao preparar query de update:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(nome, descricao, preco, quantidade, id)
	if err != nil {
		log.Println("Erro ao executar update:", err)
		return
	}

	log.Printf("Produto atualizado: ID %d - %s", id, nome)
}
