package http_response

import "github.com/gin-gonic/gin"

func Success(data map[string]interface{}) map[string]interface{} {
    return gin.H{
        "success": true,
        "data": data,
    }
}

func Error(message string) map[string]interface{} {
    return gin.H{
        "success": false,
        "message": message,
    }
}
