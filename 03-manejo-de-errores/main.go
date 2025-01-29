package main

import (
	"log"
	"os"
)

// func divide(dividendo, divisor int) (int, error) {
// 	if divisor == 0 {
// 		// return 0, errors.New("No se puede dividir por cero")
// 		return 0, fmt.Errorf("No se puede dividir por cero")
// 	}
// 	return dividendo / divisor, nil
// }

// func divide(dividendo, divisor int) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println(r)
// 		}
// 	}()

// 	validateZero(divisor)
// 	fmt.Println(dividendo / divisor)
// }

// func validateZero(divisor int) {
// 	if divisor == 0 {
// 		panic("No se puede dividar por cero, entramos en panico")
// 	}
// }

// func panico() {
// 	divide(100, 10)
// 	divide(200, 25)
// 	divide(34, 0)
// 	divide(100, 4)
// }

func main() {

	// resultado, error := divide(10, 0)

	// if error != nil {
	// 	fmt.Println("Ocurrio un error ", error)
	// 	return
	// }

	// fmt.Println("Resultado: ", resultado)

	// defer fmt.Println(3)
	// fmt.Println(1)
	// fmt.Println(2)

	// panico()
	// str := "123F"
	// num, err := strconv.Atoi(str)

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	// fmt.Println("Numero: ", num)
	// Esto agrega un prefijo a los logs
	log.SetPrefix("main(): ")
	// log.Print("Este es un super mensaje dentro de un log")

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Oye, soy un log")
}
