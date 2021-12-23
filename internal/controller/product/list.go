package productController

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "workplace/internal/database"
    "workplace/internal/dto"
    productDto "workplace/internal/dto/product"
    "workplace/internal/http/response"
    "workplace/internal/injector"
)

func List(context *gin.Context) {
    var pager dto.Pager
    if err := context.ShouldBindQuery(&pager); err != nil {
        context.JSON(http.StatusBadRequest, httpResponse.Error(err.Error()))
        return
    }

    var response productDto.Products
    response.Page = pager.Page
    response.PerPage = pager.PerPage
    injector.GetContainer().Invoke(func(repository database.ProductRepository) {
        offset := pager.PerPage * (pager.Page - 1)
        products, total, _ := repository.GetProducts(pager.PerPage, offset)
        response.Total = total
        response.Products = make([]productDto.Product, len(products))
        for key, product := range products {
            response.Products[key] = productDto.Product{
                Id: product.ID,
                Name: product.Name,
                Price: product.Price,
                UserRefer: product.UserRefer,
            }
        }
    })

    context.JSON(http.StatusOK, httpResponse.Success(response))
}
