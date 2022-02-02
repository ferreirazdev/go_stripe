package models

import (
	"database/sql"
	"time"
)

type DBModels struct {
	DB *sql.DB
}

type Models struct {
	DB DBModels
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModels{DB: db},
	}
}

type Widget struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryPrice int       `json:"inventory_price"`
	Price          int       `json:"price"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
