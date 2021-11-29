package database

import (
    "gorm.io/gorm"
    "workplace/internal/entity"
)

func Migrate(db *gorm.DB)  {
    err := db.AutoMigrate(&entity.Product{}, &entity.GrpcLog{})
    if err != nil {
        panic("failed to migration database")
    }
}
