package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	// var myBook = book.Book{
	// 	Title:  "The Amazing Knox ",
	// 	Author: "Fort Knox",
	// 	Pages:  500,
	// }

	myBook := book.NewBook("The Amazing Knox", "Knox", 500)
	// myBook.PrintInfo()

	myBook.SetTitle("The Amazing Knox (Special Edition)")
	fmt.Println(myBook.GetTitle())

	fmt.Println("********************************************")

	myBook2 := book.NewTextBook("The Amazing Knox", "Knox", 500, "Editorial Fort Knox", "Avanzado")
	// myBook2.PrintInfo()

	// myBook.PrintInfo()
	// myBook2.PrintInfo()

	book.Print(myBook)
	book.Print(myBook2)

	fmt.Println("********************************************")

	perro := animal.Perro{Nombre: "Frida "}
	gato := animal.Gato{Nombre: "Enola "}

	animal.HacerSonido(&perro)
	animal.HacerSonido(&gato)

	fmt.Println("********************************************")

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Frida2 "},
		&animal.Gato{Nombre: "Enola3 "},
		&animal.Perro{Nombre: "Frida4 "},
		&animal.Gato{Nombre: "Enola5 "},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
