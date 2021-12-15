package controller_product

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "workplace/internal/dto"
    "workplace/internal/entity"
    httpResponse "workplace/internal/http/response"
    "workplace/internal/injector"
)

func Create(context *gin.Context) {
    var reqProduct dto.Product
    if err := context.ShouldBind(&reqProduct); err != nil {
        context.JSON(http.StatusBadRequest, httpResponse.Error(err.Error()))
        return
    }
    
    var dbError error
    var dbProduct entity.Product
    injector.GetContainer().Invoke(func(db *gorm.DB) {
        dbProduct = entity.Product{
            Name: reqProduct.Name,
            Price: reqProduct.Price,
            UserRefer: reqProduct.UserRefer,
        }
        dbError = db.Create(&dbProduct).Error
    })

    if dbError != nil {
        context.JSON(http.StatusBadRequest, httpResponse.Error(dbError.Error()))
        return
    }

    context.JSON(http.StatusOK, httpResponse.Success(gin.H{
        "product_id": dbProduct.ID,
    }))
}
