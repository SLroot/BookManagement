package main

import (
	"fmt"

)

type Library struct {
	Title       string
	Author       string
	ReleaseDate int
	ID          int
}

const maxBooks = 100

var bookList [maxBooks]Library

//---------------------------------------------------------------------
// Main Function

func main() {
	fmt.Println("Welcome to book manager app!!")
	for {
		var choice int
		greeting()
		fmt.Scan(&choice)

		if choice == 1 {
			addBook(getUserStruct())
		} else if choice == 2 {
			listAllBooks()
		} else if choice == 3 {
			searchBookByID()
		} else if choice == 4 {
			deleteBookByID()
		} else {
			return
		}
	}

}

//---------------------------------------------------------------------
//CRUD Functions

func addBook(book Library) {
	for i := 0; i < maxBooks; i++ {
		if bookList[i].Title == "" {
			book.ID = i + 1
			bookList[i] = book		
			break
		}
	}
	fmt.Println()
}


func listAllBooks(){
	fmt.Println()
	var existBook []Library
	for i := 0; i < len(bookList);  i++ {
		if bookList[i].Title != "" {
			existBook = append(existBook, bookList[i])
		}
	}
	for _, v := range existBook {
		fmt.Println(v)
	}
	fmt.Println()
}

func searchBookByID() {
	fmt.Println()
	var id int
	var found bool = false

	fmt.Print("enter your id: ")
	fmt.Scan(&id)
	for i, _ := range bookList {
		if id == bookList[i].ID {
			myValue := bookList[i]
			fmt.Println(myValue)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("book not found!!")
	}
	fmt.Println()
}

func deleteBookByID() {
	fmt.Println()
	var id int
	fmt.Print("enter your id: ")
	fmt.Scan(&id)
	for i, _ := range bookList{
		if id == bookList[i].ID {
			bookList[i] = Library{}
			fmt.Println("deleted successfuly")
		} else {
			fmt.Println("book not found")
			break
		}
	}
	fmt.Println()
}


//--------------------------------------------------------------------
// Utility Functions

func getUserStruct() Library {
	var title, author string
	var realase int

	fmt.Print("enter your Title: ")
	fmt.Scan(&title)

	fmt.Print("enter your Author: ")
	fmt.Scan(&author)

	fmt.Print("enter your release: ")
	fmt.Scan(&realase)


	book := Library{
		Title: title,
		Author: author,
		ReleaseDate: realase,
	}

	return book
}

func greeting() {
	fmt.Println("1. Add Book")
	fmt.Println("2. List All Book")
	fmt.Println("3. Search Book")
	fmt.Println("4. Delete Book")
	fmt.Println("5. Exit")
	fmt.Print("What operation do you want do: ")

}