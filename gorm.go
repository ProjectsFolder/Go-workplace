package main

import (
    "fmt"
    "gorm.io/gorm"
    "workplace/internal/database"
    "workplace/internal/injector"
)

func main() {
    container := injector.GetContainer()

    err := container.Invoke(func(db *gorm.DB, repository database.ProductRepository) {
        database.Migrate(db)

        products, err := repository.GetProductsLikeName("computer2%")
        if err != nil {
            panic(err)
        }
        for _, product := range products {
            creator := "none"
            if product.Creator != nil {
                creator = product.Creator.Name
            }
            fmt.Println(fmt.Sprintf("id: %d, name: %s, user: %s",
                product.ID,
                product.Name,
                creator,
            ))
        }
    })

    if err != nil {
        panic(err)
    }
}
