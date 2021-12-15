package controller_product

import (
    "errors"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
    "workplace/internal/entity"
    httpResponse "workplace/internal/http/response"
    "workplace/internal/injector"
)

func Delete(context *gin.Context) {
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
    injector.GetContainer().Invoke(func(db *gorm.DB) {
        var dbProduct entity.Product
        db.First(&dbProduct, id)
        if (entity.Product{}) != dbProduct {
            dbError = db.Delete(&dbProduct).Error
        } else {
            dbError = errors.New("product not found")
        }
    })

    if dbError != nil {
        context.JSON(http.StatusBadRequest, httpResponse.Error(dbError.Error()))
        return
    }

    context.JSON(http.StatusOK, httpResponse.Success(struct{}{}))
}
