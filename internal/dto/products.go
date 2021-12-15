package dto

type Products struct {
    Products []Product
    Page     int
    PerPage  int
    Total    int64
}
