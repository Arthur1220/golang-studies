package main

import (
	"log"
	"personalidades/database"
	"personalidades/models"
	"personalidades/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Albert Einstein", Historia: "Físico teórico conocido por la teoría de la relatividad."},
		{Id: 2, Nome: "Marie Curie", Historia: "Pionera en el campo de la radiactividad y primera persona en ganar dos premios Nobel."},
		{Id: 3, Nome: "Isaac Newton", Historia: "Matemático y físico que formuló las leyes del movimiento y la gravitación universal."},
	}

	database.InitDatabase()

	log.Println("Iniciando sistema de Personalidades...")
	routes.HandleRequests()
}
