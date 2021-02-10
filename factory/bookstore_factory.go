package factory

import (
	"bookstore/book"
	"bookstore/suppliers"
)

type AbstractBookstoreFactory interface {
	CreateBook(string, string)
	CreateBookSupplier() suppliers.BookSupplier
}

type AudioBookFactory struct {
}

func (audioFactory AudioBookFactory) CreateBook(title, publisher string) *book.AudioBook {
	return book.NewAudioBook(title, publisher)
}

type PaperBackFactory struct {
}

func (paperback PaperBackFactory) CreateBook(title, publisher string) *book.AudioBook {
	return book.NewPaperBack(title, publisher)
}

type EBookFactory struct {
}

func (ebook EBookFactory) CreateBook(title, publisher string) *book.AudioBook {
	return book.NewEBook(title, publisher)
}
