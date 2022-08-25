package main

import (
	"fmt"
	"math"
)

func Factorial(n float64) (result float64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	calcul()
}

func calcul() {
	var a, b, res float64
	var c string

	fmt.Println("Введите число a")
	fmt.Scan(&a)

	fmt.Println("Введите число b")
	fmt.Scan(&b)

	fmt.Println("Введите оператор + - * / ^ !:")
	fmt.Scan(&c)

	switch c {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль запрещено")
			break
		}

		res = a / b
	case "^":
		res = (math.Pow(a, b))
	case "!:":
		res = Factorial(float64(a))

	}

	fmt.Printf("Результат выполнения операции: %.2f", res)
}
