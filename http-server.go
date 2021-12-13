package main

import (
    "github.com/gin-gonic/gin"
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "log"
    "time"
    "workplace/internal/config"
    controllertest "workplace/internal/controller/test"
)

func main() {
    cfg := config.GetConfig()
    rl, err := rotatelogs.New(
        cfg.RotateLogPath,
        rotatelogs.WithRotationTime(24 * time.Hour),
        rotatelogs.WithMaxAge(-1),
        rotatelogs.WithRotationCount(30),
    )
    if err != nil {
        log.Fatal("Cannot start logger:", err)
    }
    log.SetOutput(rl)
    gin.DefaultWriter = rl

    router := gin.Default()
    router.GET("/test/*id", controllertest.Test)

    err = router.Run(":" + cfg.HttpPort)
    if err != nil {
        log.Fatal("Unable to start server:", err)
    }
}
