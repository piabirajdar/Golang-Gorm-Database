package main

import (
	"fmt"

	"github.com/piabirajdar/Golang-Gorm-Database/dbaccess"
	"github.com/piabirajdar/Golang-Gorm-Database/types"
)

func genre() {
	db := dbaccess.InitDB()
	db.AutoMigrate(&types.Author{}, &types.Genre{})

	author := types.Author{Name: "Agatha Christie", Year: 2000}
	db.Create(&author)

	db.Create(&types.Genre{AuthorId: author.ID, BookName: "Sky full of stars", Genre: "Horror"})
	db.Create(&types.Genre{AuthorId: author.ID, BookName: "Think Positive", Genre: "Thriller"})

	author = types.Author{}
	authorQuery := types.Author{Name: "Agatha Christie"}
	db.Model(types.Author{}).Preload("Genres").Where(&authorQuery).First(&author)

	for _, genre := range author.Genres {
		fmt.Printf("%s: %s\n", genre.Genre, genre.BookName)
	}
}
