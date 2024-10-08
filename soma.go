package main

import "fmt"

func main() {
	sum := sum(3, 3, 3)
	multiplication := multiplication(3, 3, 3)
	division := division(4, 2)
	subtraction := subtraction(7, 1)
	fmt.Println(subtraction, division, sum, multiplication)
}

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func multiplication(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}
func subtraction(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}
func division(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		total /= v
	}
	return total
}

// func div(i ...int) int {
// 	total := 0
// 	for _, v := range i {
// 		total /= v
// 		fmt.Println(total)
// 	}
// 	return total
// }
