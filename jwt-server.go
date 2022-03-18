package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "workplace/internal/config"
    httpJwt "workplace/internal/http/jwt"
    httpMiddleware "workplace/internal/http/middleware"
    httpResponse "workplace/internal/http/response"
    "workplace/internal/injector"
)

func main() {
    var cfg *config.Configuration
    var jwtLocal *httpJwt.Jwt
    _ = injector.GetContainer().Invoke(func(config *config.Configuration, j *httpJwt.Jwt) {
        cfg = config
        jwtLocal = j
    })

    router := gin.Default()
    v1 := router.Group("/v1")
    {
        v1.GET("/login", func(context *gin.Context) {
            login := context.Query("login")
            password := context.Query("password")
            if login == "admin" && password == "test" {
                tokenString, err := jwtLocal.CreateTokenByUser(login)
                if err != nil {
                    context.JSON(http.StatusBadRequest, httpResponse.Error(err.Error()))
                    return
                }

                context.JSON(http.StatusOK, httpResponse.Success(tokenString))
                return
            }

            context.JSON(http.StatusBadRequest, httpResponse.Error("User not found"))
        })
        auth := v1.Group("/auth", httpMiddleware.JWTTokenRequiredMiddleware())
        {
            auth.GET("/hello", func(context *gin.Context) {
                claims, err := jwtLocal.GetClaimsByContext(context)
                if err != nil {
                    context.JSON(http.StatusBadRequest, httpResponse.Error("Invalid token"))
                    return
                }

                message := fmt.Sprintf("Hello, %s!", claims["user"])
                context.JSON(http.StatusOK, httpResponse.Success(message))
                return
            })
        }
    }

    err := router.Run(cfg.HttpHost + ":" + cfg.HttpPort)
    if err != nil {
        log.Fatal("Unable to start server:", err)
    }
}
