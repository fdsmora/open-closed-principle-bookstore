package book

import (
	"bookstore/misc"
)

type Book interface {
	Id() misc.ID
	Title() string
	Publisher() string
}

type book struct {
	Id               misc.ID
	Title, Publisher string
}

type AudioBook struct {
	book
	audio_format            string
	volume, player_position int
}

func NewAudioBook(title, publisher string) *AudioBook {
	ab := new(AudioBook)
	ab.book = book{misc.IdCounter.Next(), title, publisher}
	ab.volume = 10
	ab.player_position = 0
	return ab
}

func (ab *AudioBook) Id() misc.ID {
	return ab.book.Id
}
func (ab *AudioBook) Title() string {
	return ab.book.Title
}
func (ab *AudioBook) Publisher() string {
	return ab.book.Publisher
}

type PaperBack struct {
	book
	pages int
}

func NewPaperBack(title, publisher string) *PaperBack {
	pb := new(PaperBack)
	pb.book = book{misc.IdCounter.Next(), title, publisher}
	pb.pages = 300
	return pb
}

func (pb *PaperBack) Id() misc.ID {
	return pb.book.Id
}
func (pb *PaperBack) Title() string {
	return pb.book.Title
}
func (pb *PaperBack) Publisher() string {
	return pb.book.Publisher
}

type EBook struct {
	book
	pages  int
	Format string
}

func NewEBook(title, publisher string) *EBook {
	eb := new(EBook)
	eb.book = book{misc.IdCounter.Next(), title, publisher}
	eb.Format = "pdf"
	return eb
}

func (eb *EBook) Id() misc.ID {
	return eb.book.Id
}
func (eb *EBook) Title() string {
	return eb.book.Title
}
func (eb *EBook) Publisher() string {
	return eb.book.Publisher
}
