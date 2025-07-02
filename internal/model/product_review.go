package model

// import "gorm.io/gorm"

type ProductReview struct {
	BaseModel
	ProductID uint   `gorm:"column:product_id;index" json:"productId"` // Index cho truy vấn đánh giá theo sản phẩm
	UserID    uint   `gorm:"column:user_id;index" json:"userId"`       // Index cho truy vấn đánh giá theo user
	Rating    int    `gorm:"column:rating;index" json:"rating"`        // Index cho lọc/tính trung bình rating
	Comment   string `gorm:"column:comment" json:"comment"`
}
