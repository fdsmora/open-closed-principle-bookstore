package main

import (
	"bookstore/book"
	"bookstore/bookFormats"
)

func main() {
	books := []book.Book{
		{
			Id:        1,
			Title:     "El Principito",
			Publisher: "Porrua",
			Format:    bookFormats.PAPERBACK,
		},
		{
			Id:        2,
			Title:     "El amor en los tiempos del colera",
			Publisher: "Diana",
			Format:    bookFormats.AUDIOBOOK,
		},
		{
			Id:        3,
			Title:     "50 sombras de Grey",
			Publisher: "FCE",
			Format:    bookFormats.EBOOK,
		},
	}

	for _, book := range books {
		switch book.Format {
		case bookFormats.AUDIOBOOK:
			book.PlaceAudioBookOrder()
		case bookFormats.PAPERBACK:
			book.PlacePaperbackOrder()
		case bookFormats.EBOOK:
			book.PlaceKindleOrder()
		}
	}
}
