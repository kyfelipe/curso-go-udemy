package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros
	var p *int = nil

	p = &i // pegando o endereço da variável
	*p++
	i++

	fmt.Println(p, &i, *p, i)
}
