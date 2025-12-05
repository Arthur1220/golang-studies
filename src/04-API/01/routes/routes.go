package routes

import (
	"log"
	"net/http"
	"personalidades/controllers"
	"personalidades/middleware"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadePorID).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriarPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("PUT")

	log.Println("🚀 Servidor rodando! Acesse: http://localhost:8000")
	http.ListenAndServe(":8000", r)
}
