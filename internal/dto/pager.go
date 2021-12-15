package dto

type Pager struct {
    Page    int `form:"page" binding:"required,min=1"`
    PerPage int `form:"per_page" binding:"required"`
}
