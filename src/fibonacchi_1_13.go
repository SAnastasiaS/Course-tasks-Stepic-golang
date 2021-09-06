//Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

package main

import "fmt"

func main() {
	var A, previous, next, counter uint32
	fmt.Scan(&A)
	previous = 1
	next = 1
	for counter = 2; A > next; counter++ {
		next, previous = next+previous, next
	}
	switch {
	case A == next:
		fmt.Println(counter)
	default:
		fmt.Println(-1)
	}
}
