package factory

import (
	"bookstore/book"
	"bookstore/suppliers"
)

type AbstractBookstoreFactory interface {
	CreateBook(string, string) *book.Book
	CreateBookSupplier() suppliers.BookSupplier
}

type AudioBookFactory struct {
}

func (audioFactory AudioBookFactory) CreateBook(title, publisher string) *book.Book {
	return book.NewAudioBook(title, publisher)
}
func (audioFactory AudioBookFactory) CreateBookSupplier() suppliers.BookSupplier {
	return suppliers.AudioBookSupplier
}

type PaperBackFactory struct {
}

func (paperback PaperBackFactory) CreateBook(title, publisher string) *book.Book {
	return book.NewPaperBack(title, publisher)
}

func (pbFactory PaperBackFactory) CreateBookSupplier() suppliers.BookSupplier {
	return suppliers.PaperBackSupplier
}

type EBookFactory struct {
}

func (ebook EBookFactory) CreateBook(title, publisher string) *book.Book {
	return book.NewEBook(title, publisher)
}

func (ebFactory EBookFactory) CreateBookSupplier() suppliers.BookSupplier {
	return suppliers.EBookSupplier
}
