package database

import (
    "database/sql"
    "gorm.io/gorm"
    "workplace/internal/entity"
)

type ProductRepository interface {
    GetProductsLikeName(name string) []entity.Product
    GetProducts(limit int, offset int) ([]entity.Product, int64)
}

type ProductRepositoryImpl struct {
    database *gorm.DB
}

func (repository *ProductRepositoryImpl) GetProductsLikeName(name string) []entity.Product {
    var products []entity.Product
    repository.
        database.
        Raw("SELECT * FROM products p WHERE p.name LIKE @name", sql.Named("name", name)).
        Scan(&products)
    
    return products
}

func (repository *ProductRepositoryImpl) GetProducts(limit int, offset int) ([]entity.Product, int64) {
    var products []entity.Product
    repository.database.Model(&entity.Product{}).Offset(offset).Limit(limit).Find(&products)

    var count int64
    repository.database.Model(&entity.Product{}).Count(&count)

    return products, count
}

func NewProductRepository(database *gorm.DB) *ProductRepositoryImpl {
    return &ProductRepositoryImpl{database: database}
}
