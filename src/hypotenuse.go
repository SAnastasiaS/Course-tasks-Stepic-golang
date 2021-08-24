//На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

package main

import (
	"errors"
	"fmt"
	"math"
)

func hypotenuse(a float64, b float64) float64 {
	if a <= 0 || b <= 0 {
		panic(errors.New("Значения катетов не положительные, проверьте входные значения"))
	}
	return math.Sqrt(a*a + b*b)
}

func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Проверьте типы входных параметров")
	} else {
		fmt.Println(hypotenuse(a, b))
	}
}
