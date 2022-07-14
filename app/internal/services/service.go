package services

import "database/sql"

type Service interface {
	Product() *Product
}

type BaseService struct {
	Service
	db *sql.DB
}

func NewBaseService(db *sql.DB) *BaseService {
	return &BaseService{
		db: db,
	}
}

func (s *BaseService) Product() *Product {
	return NewProductService(s.db)
}
