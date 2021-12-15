package dto

type Product struct {
    Name      string `form:"name" binding:"required"`
    Price     uint `form:"price" binding:"required"`
    UserRefer int `form:"user_id" binding:"required"`
}
