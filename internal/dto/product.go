package dto

type Product struct {
    Id        uint
    Name      string `form:"name" binding:"required"`
    Price     uint `form:"price" binding:"required"`
    UserRefer int `form:"user_id" binding:"required"`
}
