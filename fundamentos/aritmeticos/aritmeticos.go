package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b) // 3 & 2 = 2
	fmt.Println("Ou =>", a|b) // 3 & 2 = 3
	fmt.Println("Xor =>", a^b) // 3 & 2 = 1

	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Maior =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}