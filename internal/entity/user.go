package entity

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name     string
    IsAdmin  bool `gorm:"default:false"`
    IsActive bool `gorm:"default:true"`
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
    tx.Model(&u).Update("IsActive", false)
    
    return
}
