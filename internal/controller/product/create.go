package productController

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "workplace/internal/dto/product"
    "workplace/internal/entity"
    "workplace/internal/http/response"
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

    context.JSON(http.StatusOK, httpResponse.Success(dto.Product{
        Id: dbProduct.ID,
        Name: dbProduct.Name,
        Price: dbProduct.Price,
        UserRefer: dbProduct.UserRefer,
    }))
}
