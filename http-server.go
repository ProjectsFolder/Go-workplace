package main

import (
    "github.com/gin-gonic/gin"
    rotateLogs "github.com/lestrrat-go/file-rotatelogs"
    "log"
    "time"
    "workplace/internal/config"
    controllerTest "workplace/internal/controller/test"
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
    router.GET("/test/*id", controllerTest.Test)

    err = router.Run(":" + cfg.HttpPort)
    if err != nil {
        log.Fatal("Unable to start server:", err)
    }
}
