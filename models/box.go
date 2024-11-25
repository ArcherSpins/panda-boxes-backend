package models

type Box struct {
	ID    string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name  string `gorm:"size:255" json:"name"`
	Price uint   `json:"price"`
}

type NewBox struct {
	Name  string `gorm:"size:255" json:"name"`
	Price uint   `json:"price"`
}
