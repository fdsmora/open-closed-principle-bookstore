package main

import (
	"bookstore/bookFormats"
	"bookstore/orders"
)

func main() {

	booksInfo := []orders.BookInfoEntry{
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

	bookOrders := orders.GenerateOrders(booksInfo)

	for _, order := range bookOrders {
		order.BookSupplier(order.Book)
	}

}
