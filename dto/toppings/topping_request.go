package toppingsdto

type CreateToppingRequest struct {
	Title string `json:"title" form:"title" gorm:"type: varchar(255)" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
}

type UpdateToppingRequest struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
	Image string `json:"image" form:"image"`
}
