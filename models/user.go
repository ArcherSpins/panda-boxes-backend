package models

type User struct {
	ID       string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Username string `gorm:"size:255;unique" json:"username"`
	Email    string `gorm:"size:255;unique" json:"email"`
	Password string `gorm:"size:255" json:"password"`
}
