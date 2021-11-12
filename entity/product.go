package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name      string
	Price     uint
	UserRefer int
	Creator   User `gorm:"foreignKey:UserRefer"`
}
