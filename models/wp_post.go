package models

type WpPost struct {
	ID        int `gorm:"column:ID"`
	PostName  string
	PostTitle string
}
