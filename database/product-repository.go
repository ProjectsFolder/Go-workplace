package database

import (
    "database/sql"
    "workplace/entity"
)

func GetProductsLikeName(name string) []entity.Product {
    db, _ := GetConnection()
    var products []entity.Product
    db.Raw("SELECT * FROM products p WHERE p.name LIKE @name", sql.Named("name", name)).Scan(&products)
    
    return products
}
