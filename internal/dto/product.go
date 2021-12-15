package dto

type Product struct {
    Id        uint `json:"id"`
    Name      string `form:"name" json:"name" binding:"required"`
    Price     uint `form:"price" json:"price" binding:"required"`
    UserRefer int `form:"user_id" json:"user_id" binding:"required"`
}
