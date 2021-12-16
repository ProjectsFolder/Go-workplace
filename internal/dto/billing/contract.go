package dto

type Contract struct {
    Id      int `json:"id"`
    Title   string `json:"title"`
    Balance float32 `json:"balance"`
    HouseId int `json:"house_id,string"`
}
