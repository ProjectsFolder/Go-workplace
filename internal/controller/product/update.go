package controller_product

import (
    "errors"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
    "workplace/internal/dto"
    "workplace/internal/entity"
    httpResponse "workplace/internal/http/response"
    "workplace/internal/injector"
)

func Update(context *gin.Context) {
    var reqProduct dto.Product
    if err := context.ShouldBind(&reqProduct); err != nil {
        context.JSON(http.StatusBadRequest, httpResponse.Error(err.Error()))
        return
    }

    id, err := strconv.Atoi(context.Param("id"))
    if err != nil {
        context.JSON(http.StatusBadRequest, httpResponse.Error("product id is not integer"))
        return
    }
    if id == 0 {
        context.JSON(http.StatusBadRequest, httpResponse.Error("product id is required"))
        return
    }

    var dbError error
    dbProduct := entity.Product{}
    injector.GetContainer().Invoke(func(db *gorm.DB) {
        db.First(&dbProduct, id)
        if (entity.Product{}) != dbProduct {
            dbError = db.Model(&dbProduct).Updates(entity.Product{
                Name: reqProduct.Name,
                Price: reqProduct.Price,
                UserRefer: reqProduct.UserRefer,
            }).Error
        } else {
            dbError = errors.New("product not found")
        }
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
