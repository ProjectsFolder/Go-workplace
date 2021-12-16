package dto

import "workplace/internal/dto"

type ProductsResponse struct {
    Data Products `json:"data"`
    dto.Response
}
