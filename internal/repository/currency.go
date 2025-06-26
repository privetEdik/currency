package repository

import (
	"currency/internal/model"
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]model.Currency, error) {
	rows, err := r.db.Query("SELECT id, name, code, sign FROM currencies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []model.Currency
	for rows.Next() {
		var c model.Currency
		if err := rows.Scan(&c.ID, &c.Name, &c.Code, &c.Sign); err != nil {
			return nil, err
		}
		currencies = append(currencies, c)
	}
	return currencies, nil
}

func (r *Repository) GetByCode(code string) (*model.Currency, error) {
	row := r.db.QueryRow("SELECT id, name, code, sign FROM currencies WHERE code = $1", code)

	var c model.Currency
	if err := row.Scan(&c.ID, &c.Name, &c.Code, &c.Sign); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}
func (r *Repository) Insert(c model.Currency) error {
	_, err := r.db.Exec("INSERT INTO currencies (name, code, sign) VALUES ($1, $2, $3)", c.Name, c.Code, c.Sign)
	return err
}
