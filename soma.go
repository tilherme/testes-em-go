package main

import "fmt"

func main() {
	x := soma(3, 3, 3)
	y := multiplica(3, 3, 3)
	fmt.Println(x, y)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}
