package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"personalidades/database"
	"personalidades/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessou a rota /")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessou a rota /personalidades")

	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func PersonalidadePorID(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessou a rota /personalidades/{id}")
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)

}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessou a rota /personalidades (criar)")

	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessou a rota /personalidades/{id} (deletar)")
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessou a rota /personalidades/{id} (editar)")
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
