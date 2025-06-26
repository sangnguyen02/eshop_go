package model

// "gorm.io/gorm"

type User struct {
	BaseModel
	Username        string          `gorm:"column:username;type:VARCHAR(100);uniqueIndex" json:"username"`
	UserCredentials UserCredentials `gorm:"foreignKey:UserID" json:"userCredentials"`
	FullName        string          `gorm:"column:fullname" json:"fullname"`
	Phone           string          `gorm:"column:phone;type:VARCHAR(20);uniqueIndex" json:"phone"`
	Email           string          `gorm:"column:email;type:VARCHAR(255);uniqueIndex" json:"email"`
	Role            string          `gorm:"column:role" json:"role"`
	Status          bool            `gorm:"column:status;index" json:"status"`
}

type UserCredentials struct {
	BaseModel
	Password string `gorm:"column:password" json:"password"`
	UserID   uint   `gorm:"column:user_id;index" json:"userId"`
}
