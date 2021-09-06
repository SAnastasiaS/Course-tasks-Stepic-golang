//Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1. Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... Напишите функцию, которая по указанному натуральному n возвращает φn.

package main

import "fmt"

func fibonacci(n int) int {
	var previous, next, counter int
	previous = 1
	next = 1
	for counter = 2; counter < n; counter++ {
		next, previous = next+previous, next
	}
	switch {
	case n == 0 || n == 1:
		return 1
	default:
		return next
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Число Фибоначчи, соответствующее номеру %d: %d \n", n, fibonacci(n))
}
