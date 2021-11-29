package database

import (
    "database/sql"
    "gorm.io/gorm"
    "workplace/internal/entity"
)

type ProductRepository struct {
    database *gorm.DB
}

func (repository *ProductRepository) GetProductsLikeName(name string) []entity.Product {
    var products []entity.Product
    repository.
        database.
        Raw("SELECT * FROM products p WHERE p.name LIKE @name", sql.Named("name", name)).
        Scan(&products)
    
    return products
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
    return &ProductRepository{database: database}
}
