package main

import (
	"ValidacaoTesteMais/database"
	"ValidacaoTesteMais/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
