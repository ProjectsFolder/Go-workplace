package dto

type Contract struct {
    Id      int `json:"id"`
    Title   string `json:"title"`
    Balance float32 `json:"balance"`
    HouseId string `json:"house_id"`
}
