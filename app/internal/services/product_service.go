package services

import (
	"database/sql"
	"fmt"
	"product-service/internal/model"
	"time"
)

type Product struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) *Product {
	return &Product{
		db: db,
	}
}

func (p *Product) Create(product model.CreateProductDTO) (sql.Result, error) {
	query := fmt.Sprintf(
		CREATE_PRODUCT,
		product.Name,
		product.Description,
		product.Price,
		product.SpecToJson(),
		product.CategoryId,
		time.Now(),
		time.Now())
	res, err := p.db.Exec(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *Product) Delete(id int) (sql.Result, error) {
	query := fmt.Sprintf(DELETE_PRODUCT, id)
	res, err := p.db.Exec(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *Product) Update(product model.UpdateProductDTO) (sql.Result, error) {
	query := fmt.Sprintf(
		UPDATE_PRODUCT,
		product.Name,
		product.Description,
		product.Price,
		product.SpecToJson(),
		product.CategoryId,
		time.Now(),
		product.Rating,
		product.Id)
	res, err := p.db.Exec(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *Product) FetchOne(id int) (*model.ProductDB, error) {
	query := fmt.Sprintf(SELECT_PRODUCT, id)
	res, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	products, err := p.Scan(res)
	if err != nil {
		return nil, err
	}
	return &products[0], nil
}

func (p *Product) FetchAll() ([]model.ProductDB, error) {
	query := fmt.Sprintf(SELECT_ALL_PRODUCT)
	res, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	products, err := p.Scan(res)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (product *Product) Scan(rows *sql.Rows) ([]model.ProductDB, error) {
	var products []model.ProductDB
	for rows.Next() {
		var product model.ProductDB
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.ImageId,
			&product.Price,
			&product.CurrencyId,
			&product.Rating,
			&product.Specification,
			&product.CategoryId,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
