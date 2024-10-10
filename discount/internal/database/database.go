package database

import (
	"database/sql"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
}

type service struct {
	db *sql.DB
}

func New() Service {
	db, err := sql.Open("sqlite3", "../../discount.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS Discount (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "product_name" TEXT,
		"description" TEXT,
        "amount" INTEGER
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err != nil {
		log.Fatal(err)

	}
	return &service{
		db: db,
	}
}

func (s *service) Health() map[string]string {
	err := s.db.Ping()
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
