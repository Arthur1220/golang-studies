package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	// Em produção, isso viria de os.Getenv("DATABASE_URL")
	conexao := "user=postgres dbname=alura_loja password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		// Aqui panic é aceitável pois sem banco a app não inicia
		log.Fatal("Erro ao conectar driver:", err)
	}

	// Validar se a conexão está realmente ativa
	if err = db.Ping(); err != nil {
		log.Fatal("Erro ao pingar banco:", err)
	}

	return db
}
