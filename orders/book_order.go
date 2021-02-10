package orders

import (
	"bookstore/book"
	"bookstore/bookFormats"
	"bookstore/factory"
	"bookstore/suppliers"
)

type BookInfoEntry map[string]string

type BookOrder struct {
	Book         *book.Book
	BookSupplier suppliers.BookSupplier
}

func NewBookOrder(bookData BookInfoEntry, bookstoreFactory factory.AbstractBookstoreFactory) *BookOrder {
	book := bookstoreFactory.CreateBook(bookData["Title"], bookData["Publisher"])
	bookSuplier := bookstoreFactory.CreateBookSupplier()
	return &BookOrder{book, bookSuplier}
}

func GenerateOrders(bookInfo []BookInfoEntry) []*BookOrder {
	orders := make([]*BookOrder, len(bookInfo))
	var i = 0

	for _, bookData := range bookInfo {
		orders[i] = CreateOrder(bookData)
		i++
	}
	return orders
}

func CreateOrder(bookData BookInfoEntry) *BookOrder {
	switch bookData["Format"] {
	case bookFormats.AUDIOBOOK:
		return NewBookOrder(bookData, factory.AudioBookFactory{})
	case bookFormats.PAPERBACK:
		return NewBookOrder(bookData, factory.PaperBackFactory{})
	case bookFormats.EBOOK:
		return NewBookOrder(bookData, factory.EBookFactory{})
	}
	return nil
}
