package book

import (
	"bookstore/bookFormats"
	"bookstore/factory"
	"bookstore/misc"
	"bookstore/suppliers"
)

type BookInfoEntry map[string]string

type BookOrder struct {
	book         Book
	bookSupplier suppliers.BookSupplier
}

func NewBookOrder(bookData BookInfoEntry, bookstoreFactory factory.AbstractBookstoreFactory) *BookOrder {
	book := bookstoreFactory.CreateBook(bookData["Title"], bookData["Publisher"])
	bookSuplier := bookstoreFactory.CreateBookSupplier()
	return &BookOrder{book, bookSuplier}
}

func GenerateOrders(bookInfo []BookInfoEntry) {
	orders := make([]*BookOrder, len(bookInfo))
	var i = 0

	for _, bookData := range bookInfo {
		switch bookData["Format"] {
		case bookFormats.AUDIOBOOK:
			orders[i] = NewBookOrder(bookData, factory.AudioBookFactory{})
		case bookFormats.PAPERBACK:
			orders[i] = NewBookOrder(bookData, factory.PaperBackFactory{})
		case bookFormats.EBOOK:
			orders[i] = NewBookOrder(bookData, factory.EBookFactory{})
		}
		i++
	}
	return orders
}

type Book struct {
	Id               misc.ID
	Title, Publisher string
}

type AudioBook struct {
	Book
	audio_format            string
	volume, player_position int
}

func NewAudioBook(title, publisher string) *AudioBook {
	ab := new(AudioBook)
	ab.Book = Book{misc.IdCounter.Next(), title, publisher}
	ab.volume = 10
	ab.player_position = 0
	return ab
}

type PaperBack struct {
	Book
	pages int
}

func NewPaperBack(title, publisher string) *PaperBack {
	pb := new(PaperBack)
	pb.Book = Book{misc.IdCounter.Next(), title, publisher}
	pb.pages = 300
	return pb
}

type EBook struct {
	Book
	pages  int
	Format string
}

func NewEBook(title, publisher string) *EBook {
	eb := new(EBook)
	eb.Book = Book{misc.IdCounter.Next(), title, publisher}
	eb.Format = "pdf"
	return eb
}
