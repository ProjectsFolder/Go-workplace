package httpJwt

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
    "regexp"
    "time"
)

type Jwt struct {
    secret []byte
}

func NewJWT(secret string) *Jwt {
    return &Jwt{
        secret: []byte(secret),
    }
}

func (j *Jwt) CreateTokenByUser(login string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user": login,
        "exp": time.Now().Add(time.Minute * 15).Unix(),
    })
    tokenString, err := token.SignedString(j.secret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (j *Jwt) ValidateToken(tokenString string) (bool, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return j.secret, nil
    })

    if err != nil {
        return false, err
    }

    if !token.Valid {
        return false, err
    }

    return true, nil
}

func (j *Jwt) GetClaimsByContext(context *gin.Context) (jwt.MapClaims, error) {
    tokenString := context.GetHeader("Authorization")
    re := regexp.MustCompile(`^Bearer\s`)
    tokenString = re.ReplaceAllString(tokenString, "")
    token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return j.secret, nil
    })
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        return claims, nil
    }

    return nil, fmt.Errorf("error")
}
