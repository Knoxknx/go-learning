package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Numbers interface {
	~int | ~float64 | ~uint | ~float32
}

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	// PrintList("Knox", "Knox 2", 23, 0.14)
	// PrintList2("Knox", "Knox 2", 23, 0.14)

	fmt.Println(Sum(4.6, 5, 6, 4))

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4}

	fmt.Println(includes(strings, "a"))
	fmt.Println(includes(numbers, 5))

	fmt.Println(Filter(numbers, func(value int) bool {
		return value < 3
	}))

	product1 := Product[uint]{1, "tillas locas", 70}
	product2 := Product[string]{"U1B-SOD-921", "tillas locas", 80}

	fmt.Println(product1)
	fmt.Println(product2)
}

// * Funcion generica
func PrintList(list ...interface{}) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func PrintList2(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// func Sum(nums ...int) int {
// 	var total int
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

func includes[T comparable](list []T, valor T) bool {
	for _, item := range list {
		if item == valor {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}
