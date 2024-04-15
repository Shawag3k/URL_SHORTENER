// database.go
package main

import (
	"database/sql"
)

// Database representa a instância do banco de dados PostgreSQL.
type Database struct {
	db *sql.DB
}

// NewDatabase cria uma nova instância de Database.
func NewDatabase(connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

// SaveURL salva uma URL encurtada e sua URL original no banco de dados.
func (db *Database) SaveURL(shortURL, originalURL string) error {
	_, err := db.db.Exec("INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", shortURL, originalURL)
	if err != nil {
		return err
	}
	return nil
}

// GetOriginalURL recupera a URL original correspondente a uma URL encurtada do banco de dados.
func (db *Database) GetOriginalURL(shortURL string) (string, error) {
	var originalURL string
	err := db.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
