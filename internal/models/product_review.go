package models

// ProductReview represents a review of a product
type ProductReview struct {
	BaseModel
	ProductID uint   `gorm:"column:product_id" json:"productId"`
	UserID    uint   `gorm:"column:user_id" json:"userId"`
	Rating    int    `gorm:"column:rating" json:"rating"`
	Comment   string `gorm:"column:comment" json:"comment"`
}
