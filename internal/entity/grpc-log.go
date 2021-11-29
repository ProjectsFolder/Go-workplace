package entity

import (
    "gorm.io/gorm"
)

type GrpcLog struct {
    gorm.Model
    Message string
}
