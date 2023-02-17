package domain

type Cart struct {
	ID     uint `gorm:"primaryKey"`
	Total  uint `gorm:"not null;default:0"`
	UserId uint
}
