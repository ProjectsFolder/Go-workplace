package main

import (
    "fmt"
    "workplace/database"
    "workplace/entity"
)

func main() {
    db, err := database.GetConnection()
    if err != nil {
        panic("failed to connect database")
    }

    err = db.AutoMigrate(&entity.Product{})
    if err != nil {
        panic("failed to migration database")
    }

    user := entity.User{Name: "asylum29"}
    db.Create(&user)
    product := entity.Product{Name: "computer", Price: 1234, Creator: user}
    db.Create(&product)

    db.Delete(&product)
    db.Delete(&user)

    product = entity.Product{}
    name := "computer2"
    db.Preload("Creator").First(&product, "name = ?", name)
    if (entity.Product{}) == product {
        fmt.Println(fmt.Sprintf("not found product with name, %s", name))
        id := 1
        db.Preload("Creator").First(&product, id)
        if (entity.Product{}) == product {
            fmt.Println(fmt.Sprintf("not found product with id, %d", id))
        }
    }

    if (entity.Product{}) != product {
        fmt.Println(fmt.Sprintf("creator created %s", product.Creator.CreatedAt.Format("2006.01.02 15:04:05")))
        db.Model(&product).Update("Price", 2345)
        db.Model(&product).Updates(entity.Product{Price: 3456, Name: "computer2"})
    }

    products := database.GetProductsLikeName("compu%")
    for _, product := range products {
        fmt.Println(fmt.Sprintf("id: %d, name: %s", product.ID, product.Name))
    }
}
