package entities

type Slug struct {
	Slug   string `gorm:"unique; size:300"`
	PostID uint
	Post   Post
}
