package book

import "fmt"

type OrderRequester interface {
	PlacePaperbackOrder()
	PlaceKindleOrder()
	PlaceAudioBookOrder()
}

type Book struct {
	Id               int
	Title, Publisher string
	Format           int
}

func (book *Book) PlacePaperbackOrder() {
	fmt.Printf("Placed an order for the book '%s' in paperback format\n", book.Title)
}

func (book *Book) PlaceAudioBookOrder() {
	fmt.Printf("Placed an order for the book '%s' in audiobook format\n", book.Title)
}

func (book *Book) PlaceKindleOrder() {
	fmt.Printf("Placed an order for the book '%s' in ebook format\n", book.Title)
}
