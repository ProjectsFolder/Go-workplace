package controller_test

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis"
    "log"
    "net/http"
    "strings"
    "time"
    httpResponse "workplace/internal/http/response"
    "workplace/internal/injector"
    "workplace/internal/services"
)

func Test(context *gin.Context) {
    val := context.Query("value")
    id := strings.Replace(context.Param("id"), "/", "", 1)
    if len(val) == 0 {
        log.Println("required value query param")
        context.JSON(http.StatusBadRequest, httpResponse.Error("Wrong parameters"))
    } else {
        go func() {
            container := injector.GetContainer()
            container.Invoke(func(rc *redis.Client, logger *services.Logger) {
                rc.Set("http-redis", val + "-" + id, 60 * time.Second)
                logger.Log(val + "-" + id)
            })
        }()

        context.JSON(http.StatusOK, httpResponse.Success(gin.H{
            "message": fmt.Sprintf("Hello, %s-%s!", val, id),
        }))
    }
}
