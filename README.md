 # Book Management
 A simple coommand-line application for managing books, written in Go.


## Features
ADD a New book
List All books
Search for a book by ID
Delete a book by ID

## Book Structure

Each book contains:
 `Title` - the title of the book
 `Author` - the name of the author
 `ReleaseDate` - the year of release
 `ID` - a unique identifier (auto-generated)

## How to Run :
  1. Make sure you have Go installed
  2. Run the application with:
       ```bash
       go run main.go
  3.choose from the followins options:
    1. Add a book
    2. List all books
    3. Search a book by ID
    4. Delete a book by ID
    5. Exit

## Capacity
  the application supports up to 100 books(fixed-size-array)
  it is temprory. maybe it will be change.