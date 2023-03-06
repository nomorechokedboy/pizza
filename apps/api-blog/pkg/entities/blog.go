package entities

import "time"

type Blog struct {
	Id        uint
	Author_Id uint
	Parent_Id uint
	Title     string
	Content   string
	PublishAt time.Time
	CreateAt  time.Time
	UpdatedAt time.Time
	DeleteAt  time.Time
}
