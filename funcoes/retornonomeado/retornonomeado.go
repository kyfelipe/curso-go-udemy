package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	fmt.Println(trocar(2, 3))

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
