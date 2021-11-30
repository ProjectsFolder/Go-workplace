package test

import (
    "testing"
    "workplace/internal/entity"
)

func TestProductRepository(t *testing.T)  {
    repository := NewProductRepositoryMock(t)
    repository.GetProductsLikeNameMock.Set(func(name string) []entity.Product {
        var products []entity.Product
        if "exists" == name {
            for i := 0; i < 3; i++ {
                products = append(products, entity.Product{Name: name, Price: uint(i)})
            }
        }
        
        return products
    })

    products := repository.GetProductsLikeName("exists")
    if len(products) == 0 {
       t.Error("products is empty")
    } else {
        for _, product := range products {
            if "exists" != product.Name {
                t.Error("product name is incorrect")
            }
        }
    }

    products = repository.GetProductsLikeName("not-exists")
    if len(products) != 0 {
       t.Error("products is not empty")
    }
}
