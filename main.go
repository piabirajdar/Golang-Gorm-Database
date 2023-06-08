package main

import (
	"fmt"

	"github.com/piabirajdar/Golang-Gorm-Database/dbaccess"
	"github.com/piabirajdar/Golang-Gorm-Database/types"
)

func main() {
	db := dbaccess.InitDB()
	db.Create(&types.Author{Name: "Agatha Christie", Year: 2000})
	db.Create(&types.Author{Name: "Danielle Steel", Year: 2001})
	db.Create(&types.Author{Name: "	Harold Robbins", Year: 2002})

	var author types.Author
	//get the first row author
	db.First(&author, 1)
	fmt.Println("db.First(&author, 1)", db.First(&author, 1))
	fmt.Println("Author ID", author.ID)

	//clear out author instance
	author = types.Author{}
	//where clause ? first return
	db.Where("name=?", "Agatha Christie").First(&author)
	fmt.Printf("%v (%v)", author.Name, author.Year)

	// Update existing row
	author = types.Author{}
	authorQuery := types.Author{Name: "Agatha Christie"}

	db.Where(authorQuery).First(&author)
	db.Model(&author).Update("Year", 2009)
	fmt.Println("Updated Author Year")
	fmt.Printf("%v (%v)", author.Name, author.Year)

	//Delete (Soft delete)
	author = types.Author{}
	db.Where(&types.Author{Year: 2003}).First(&author)
	fmt.Println("To be deleted", author.Name)
	db.Delete(&author)

}
