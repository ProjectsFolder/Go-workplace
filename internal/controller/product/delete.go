package controller_product

import (
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
        dbError = db.First(&dbProduct, id).Error
        if dbError == nil {
            dbError = db.Delete(&dbProduct).Error
        }
    })

    if dbError != nil {
        context.JSON(http.StatusBadRequest, httpResponse.Error(dbError.Error()))
        return
    }

    context.JSON(http.StatusOK, httpResponse.Success(struct{}{}))
}
