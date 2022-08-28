package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	fmt.Println("Введите пять чисел:")
	fmt.Scanln(&a, &b, &c, &d, &e)
	ar := [5]int{a, b, c, d, e}
	InsertionSort(ar)
}
func InsertionSort(ar [5]int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
	fmt.Println(ar)
}
