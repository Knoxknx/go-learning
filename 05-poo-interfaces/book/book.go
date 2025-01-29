package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

// * Simulación de un constructor
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Titulo: %s\nAutor: %s\nPaginas: %d\n", b.title, b.author, b.pages)
}

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title string, author string, pages int, editorial string, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Titulo: %s\nAutor: %s\nPaginas: %d\nEditorial: %s\nNivel: %s\n", b.title, b.author, b.pages, b.editorial, b.level)
}
