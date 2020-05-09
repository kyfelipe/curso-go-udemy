package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mpas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[987654321] = "Pedro"
	aprovados[987456321] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[987456321])
	delete(aprovados, 987456321)
	fmt.Println(aprovados[987456321])
}
