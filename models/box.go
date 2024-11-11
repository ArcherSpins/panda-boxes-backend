package models

type Box struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Price uint
}
