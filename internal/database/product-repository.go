package database

import (
    "database/sql"
    "gorm.io/gorm"
    "workplace/internal/entity"
)

type ProductRepository interface {
    GetProductsLikeName(name string) ([]entity.Product, error)
    GetProducts(limit int, offset int) ([]entity.Product, int64, error)
}

type ProductRepositoryImpl struct {
    database *gorm.DB
}

func (repository *ProductRepositoryImpl) GetProductsLikeName(name string) ([]entity.Product, error) {
    var products []entity.Product
    err := repository.
        database.
        Model(&entity.Product{}).
        Where("products.name LIKE @name", sql.Named("name", name)).
        Preload("Creator").
        Find(&products).
        Error

    return products, err
}

func (repository *ProductRepositoryImpl) GetProducts(limit int, offset int) ([]entity.Product, int64, error) {
    var products []entity.Product
    if err := repository.database.Model(&entity.Product{}).Offset(offset).Limit(limit).Find(&products).Error; err != nil {
        return products, 0, err
    }

    var count int64
    if err := repository.database.Model(&entity.Product{}).Count(&count).Error; err != nil {
        return products, 0, err
    }

    return products, count, nil
}

func NewProductRepository(database *gorm.DB) *ProductRepositoryImpl {
    return &ProductRepositoryImpl{database: database}
}
