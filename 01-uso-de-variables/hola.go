package main

import (
	"fmt"
)

const Pi = 3.14

const (
	x = 100
	y = 0b1010 // Binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	// var firstName, lastName string
	// var age int

	// var (
	// 	firstName, lastName string
	// 	age                 int
	// )

	// var firstName, lastName, age = "Fort", "Knox", 29

	// usando := las variables solo se puede ocupar dentro de las funciones
	// usando var las variables se puede ocupar en todas las funciones
	firstName, lastName, age := "Fort", "Knox", 29

	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
}
