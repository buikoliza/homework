package main

import (
	"fmt"
	"math"
)

func main() {
	Srect()
	Scirc()
	Num()
}
func Srect() {
	var a, b int8
	fmt.Print("Введите сторону а прямоугольника = ")
	fmt.Scan(&a)
	fmt.Print("Введите сторону b прямоугольника = ")
	fmt.Scan(&b)

	var c = (a * b)
	fmt.Println("Площадь прямоугольника равна ="+"", c)
}
func Scirc() {
	var S float64
	fmt.Print("Введите площадь окружности =")
	fmt.Scan(&S)

	var d = math.Round(2 * (math.Sqrt(S / math.Pi)))
	fmt.Println("Диаметр окружности равен ="+"", d)

	var o = math.Round(math.Sqrt(S * 4 * math.Pi))
	fmt.Println("Длина окружности равна ="+"", o)
}
func Num() {
	var i float64
	fmt.Print("Введите трехзначное число =")
	fmt.Scan(&i)

	var a = math.Floor(i / 100)
	fmt.Println("Количество сотен = ", a)

	var b = math.Floor(i/100*10 - (math.Floor(i/100) * 10))
	fmt.Println("Количество десятков =", b)

	var c = (i - (a * 100) - (b * 10))
	fmt.Println("Количество единиц =", c)
}
