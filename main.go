package main

import (
	"bookstore/book"
	"bookstore/bookFormats"
	"fmt"
)

func main() {

	booksInfo := [3]book.BookInfoEntry{
		{
			"Title":     "El Principito",
			"Publisher": "Porrua",
			"Format":    bookFormats.PAPERBACK,
		},
		{
			"Title":     "El amor en los tiempos del colera",
			"Publisher": "Diana",
			"Format":    bookFormats.AUDIOBOOK,
		},
		{
			"Title":     "50 sombras de Grey",
			"Publisher": "FCE",
			"Format":    bookFormats.EBOOK,
		},
	}

	bookOrders := book.GenerateOrders(booksInfo)
	fmt.Println(bookOrders)

}
