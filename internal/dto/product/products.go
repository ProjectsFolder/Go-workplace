package dto

type Products struct {
    Products []Product `json:"products"`
    Page     int       `json:"page"`
    PerPage  int `json:"per_page"`
    Total    int64 `json:"total"`
}
