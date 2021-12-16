package controller_product

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "workplace/internal/database"
    "workplace/internal/dto"
    httpResponse "workplace/internal/http/response"
    "workplace/internal/injector"
)

func List(context *gin.Context) {
    var pager dto.Pager
    if err := context.ShouldBindQuery(&pager); err != nil {
        context.JSON(http.StatusBadRequest, httpResponse.Error(err.Error()))
        return
    }

    var response dto.Products
    response.Page = pager.Page
    response.PerPage = pager.PerPage
    injector.GetContainer().Invoke(func(repository database.ProductRepository) {
        offset := pager.PerPage * (pager.Page - 1)
        products, total := repository.GetProducts(pager.PerPage, offset)
        response.Total = total
        response.Products = make([]dto.Product, len(products))
        for key, product := range products {
            response.Products[key] = dto.Product{
                Id: product.ID,
                Name: product.Name,
                Price: product.Price,
                UserRefer: product.UserRefer,
            }
        }
    })

    context.JSON(http.StatusOK, httpResponse.Success(response))
}
