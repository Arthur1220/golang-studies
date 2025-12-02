package main

import (
	"log"
	"net/http"

	"loja/db"
	"loja/routes"
)

func main() {
	log.Println("Iniciando sistema da Loja...")

	log.Println("Testando conexão com o Banco de Dados...")
	database := db.ConnectDatabase()

	database.Close()
	log.Println("✅ Banco de Dados conectado com sucesso!")

	routes.SetupRoutes()
	log.Println("Rotas carregadas.")

	log.Println("🚀 Servidor rodando! Acesse: http://localhost:8000")

	http.ListenAndServe(":8000", nil)
}
