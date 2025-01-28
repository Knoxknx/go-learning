package main

import (
	"fmt"
	"math"
)

func tipos() {

	var integer int = -12
	// uint no permite valores negativos
	var integer2 uint = 12
	var float float32 = 12.12

	var a byte = 'a'

	s := "hello"

	var r rune = '❤'

	fmt.Println(integer, integer2, float)
	fmt.Println(math.MaxInt64, math.MinInt64)
	fmt.Println(a)
	fmt.Println(s[0])
	fmt.Println(r)

}
