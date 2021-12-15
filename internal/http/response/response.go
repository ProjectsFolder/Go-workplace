package http_response

func Success(data interface{}) map[string]interface{} {
    return map[string]interface{}{
        "success": true,
        "data": data,
    }
}

func Error(message string) map[string]interface{} {
    return map[string]interface{}{
        "success": false,
        "message": message,
    }
}
