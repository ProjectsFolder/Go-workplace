package httpMiddleware

import (
    "github.com/gin-gonic/gin"
    "workplace/internal/config"
    "workplace/internal/http/response"
    "workplace/internal/injector"
)

func TokenRequiredMiddleware() gin.HandlerFunc {
    var token string
    injector.GetContainer().Invoke(func(cfg *config.Configuration) {
        token = cfg.ApiToken
    })
    return func(c *gin.Context) {
        qToken := c.Query("token")

        if qToken == "" {
            c.AbortWithStatusJSON(401, httpResponse.Error("API token required"))
            return
        }

        if qToken != token {
            c.AbortWithStatusJSON(401, httpResponse.Error("Invalid API token"))
            return
        }

        c.Next()
    }
}
