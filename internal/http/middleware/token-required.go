package httpMiddleware

import (
    "github.com/gin-gonic/gin"
    "regexp"
    "workplace/internal/config"
    httpJwt "workplace/internal/http/jwt"
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

func JWTTokenRequiredMiddleware() gin.HandlerFunc {
    var jLocal *httpJwt.Jwt
    injector.GetContainer().Invoke(func(j *httpJwt.Jwt) {
        jLocal = j
    })
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        re := regexp.MustCompile(`^Bearer\s`)
        tokenString = re.ReplaceAllString(tokenString, "")
        valid, err := jLocal.ValidateToken(tokenString)

        if err != nil || !valid {
            c.AbortWithStatusJSON(401, httpResponse.Error("incorrect jwt-token"))
            return
        }

        c.Next()
    }
}
