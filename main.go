package main

import (
	"currency/internal/handler"
	"currency/internal/repository"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/currency-db?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := initDB(db); err != nil {
		log.Fatal("Failed to initialize database", err)
	}

	repo := repository.NewRepository(db)
	handler := handler.NewHandler(repo)

	http.HandleFunc("/currency", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetAll(w, r)
			/*} else if r.Method == http.MethodPost {
			handler.Create(w, r)*/
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/currencies", handler.Create)
	http.HandleFunc("/currency/", handler.GetByCode)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB(db *sql.DB) error {
	files := []string{
		"db/001_create_tables.sql",
		"db/002_seed_data.sql",
	}

	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		if _, err := db.Exec(string(bytes)); err != nil {
			return err
		}
	}
	return nil
}
