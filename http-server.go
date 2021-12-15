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
        v1.POST("/product/create", productController.Create)
    }

    err = router.Run(":" + cfg.HttpPort)
    if err != nil {
        log.Fatal("Unable to start server:", err)
    }
}
