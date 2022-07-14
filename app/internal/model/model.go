package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

type ProductDB struct {
	Id            int            `json:"id" db:"id"`
	Name          string         `json:"name" db:"name"`
	Description   string         `json:"description" db:"description"`
	Price         float64        `json:"price" db:"price"`
	CreatedAt     string         `json:"created_at" db:"created_at"`
	UpdatedAt     string         `json:"updated_at" db:"updated_at"`
	Specification sql.NullString `json:"specification" db:"specification"`
	Rating        sql.NullInt32  `json:"rating" db:"rating"`
	ImageId       sql.NullString `json:"image_id" db:"image_id"`
	CategoryId    int            `json:"category_id" db:"category_id"`
	CurrencyId    sql.NullString `json:"currency_id" db:"currency_id"`
}

type Entity struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Specification string    `json:"specification"`
}

func (p *Entity) SpecToJson() json.RawMessage {
	b, _ := json.Marshal(p.Specification)
	return b
}

type CreateProductDTO struct {
	Entity
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryId  int     `json:"category_id"`
}

type UpdateProductDTO struct {
	Entity
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Price       float64        `json:"price"`
	Description string         `json:"description"`
	Rating      int            `json:"rating"`
	ImageId     sql.NullString `json:"image_id"`
	CategoryId  int            `json:"category_id"`
}
