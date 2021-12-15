package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
    "workplace/internal/config"
    "workplace/internal/dto"
)

func main() {
    cfg := config.GetConfig()
    
    client := &http.Client{
        Timeout: 30 * time.Second,
    }
    url := "http://localhost:"+cfg.HttpPort+"/v1/product"
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalln(err)
    }
    q := request.URL.Query()
    q.Add("token", cfg.ApiToken)
    q.Add("page", "1")
    q.Add("per_page", "10")
    request.URL.RawQuery = q.Encode()

    response, err := client.Do(request)
    if err != nil {
        log.Fatalln(err)
    }
    
    bodyBytes, _ := io.ReadAll(response.Body)
    var result dto.ProductsResponse
    json.Unmarshal(bodyBytes, &result)
    
    fmt.Println(result.Success)
    fmt.Println(result.Message)
    for _, product := range result.Data.Products {
        fmt.Println(fmt.Sprintf(
            "id: %d; name: %s; price: %d; user_id: %d",
            product.Id,
            product.Name,
            product.Price,
            product.UserRefer,
            ),
        )
    }
}
