package database

import (
    "workplace/entity"
)

func Migrate()  {
    db, _ := GetConnection()
    err := db.AutoMigrate(&entity.Product{})
    if err != nil {
        panic("failed to migration database")
    }
}
