package main

import (
    "github.com/gin-gonic/gin"
    rotateLogs "github.com/lestrrat-go/file-rotatelogs"
    "log"
    "time"
    "workplace/internal/config"
    productController "workplace/internal/controller/product"
    testController "workplace/internal/controller/test"
    httpMiddleware "workplace/internal/http/middleware"
)

func main() {
    cfg := config.GetConfig()
    rl, err := rotateLogs.New(
        cfg.RotateLogPath,
        rotateLogs.WithRotationTime(24 * time.Hour),
        rotateLogs.WithMaxAge(-1),
        rotateLogs.WithRotationCount(30),
    )
    if err != nil {
        log.Fatal("Cannot start logger:", err)
    }
    log.SetOutput(rl)
    gin.DefaultWriter = rl

    router := gin.Default()
    v1 := router.Group("/v1", httpMiddleware.TokenRequiredMiddleware())
    {
        v1.GET("/test/*id", testController.Test)
        product := v1.Group("/product")
        {
            product.GET("/", productController.List)
            product.GET("/:id", productController.View)
            product.POST("/", productController.Create)
            product.PUT("/:id", productController.Update)
            product.DELETE("/:id", productController.Delete)
        }
    }

    err = router.Run(":" + cfg.HttpPort)
    if err != nil {
        log.Fatal("Unable to start server:", err)
    }
}
