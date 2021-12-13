package controller_test

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis"
    "log"
    "net/http"
    "strings"
    "time"
    "workplace/internal/injector"
    "workplace/internal/services"
)

func Test(context *gin.Context) {
    val := context.Query("value")
    id := strings.Replace(context.Param("id"), "/", "", 1)
    if len(val) == 0 {
        log.Println("required value query param")
        context.String(http.StatusBadRequest, "wrong parameters")
    } else {
        go func() {
            container := injector.GetContainer()
            container.Invoke(func(rc *redis.Client, logger *services.Logger) {
                rc.Set("http-redis", val + "-" + id, 60 * time.Second)
                logger.Log(val + "-" + id)
            })
        }()

        context.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf("Hello, %s-%s!", val, id),
        })
    }
}
