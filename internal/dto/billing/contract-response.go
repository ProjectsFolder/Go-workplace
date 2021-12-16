package dto

import "workplace/internal/dto"

type ContractResponse struct {
    Data []Contract `json:"data"`
    dto.Response
}
