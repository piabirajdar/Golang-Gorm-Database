package types

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name   string `gorm:"uniqueIndex"`
	Year   uint   `gorm:"not null"`
	Genres []Genre
	//CreatedAt
	//UpdatedAt
	//DeletedAt
}

type Genre struct {
	gorm.Model
	BookName string
	Genre    string
	AuthorId uint
}
