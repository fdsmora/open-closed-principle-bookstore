package suppliers

import (
	"bookstore/book"
	"fmt"
)

type BookSupplier func(*book.Book)

var AudioBookSupplier BookSupplier = func(book *book.Book) {
	fmt.Println("Submitted order for producing the book ", book.Title, "in audiobook format")
}

var PaperBackSupplier = func(book *book.Book) {
	fmt.Println("Submitted order for producing the book ", book.Title, "in paperback format")
}

var EBookSupplier = func(book *book.Book) {
	fmt.Println("Submitted order for producing the book ", book.Title, "in ebook format")
}
