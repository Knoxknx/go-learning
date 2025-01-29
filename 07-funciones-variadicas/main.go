package main

import "fmt"

func main() {
	fmt.Println(suma("knox", 12, 45, 78, 56))
	fmt.Println(suma("knox", 10, 20, 30, 40, 50))

	imprimirDatos("hola", 28, true, 3.14)

	fmt.Println(factorial(3))

	//función anonima
	saludo := func(name string) {
		fmt.Printf("Hola, %s \n", name)
	}

	saludo("Knox")

	saludar("Knox", saludo)

	fmt.Println("************************************************************************")
	r := double(addOne, 3)
	fmt.Println("Resultado: ", r)

	fmt.Println("************************************************************************")
	nextInt := incrementer()
	fmt.Println("incrementer: ", nextInt())
	fmt.Println("incrementer: ", nextInt())
	fmt.Println("incrementer: ", nextInt())
	fmt.Println("incrementer: ", nextInt())
	fmt.Println("incrementer: ", nextInt())
}

func suma(name string, nums ...int) int {
	var total int

	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s, la suma es: %d\n", name, total)
	return total
}

// Funcion variadica
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

// Funcion recusiva
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func saludar(name string, f func(string)) {
	f(name)
}

func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

// Funcion closures
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
