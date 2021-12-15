package main

import (
    "fmt"
    "gorm.io/gorm"
    "workplace/internal/database"
    "workplace/internal/injector"
)

func main() {
    container := injector.GetContainer()

    err := container.Invoke(func(db *gorm.DB, repository *database.ProductRepositoryImpl) {
        database.Migrate(db)

        products := repository.GetProductsLikeName("compu%")
        for _, product := range products {
            fmt.Println(fmt.Sprintf("id: %d, name: %s", product.ID, product.Name))
        }
    })

    if err != nil {
        panic(err)
    }
}
