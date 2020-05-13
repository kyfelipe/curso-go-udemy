package main

import (
	"fmt"
	"github.com/kyfelipe/html"
)

func encaminhar(origem <- chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar1 - misturar (mensagens) em um único canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.google.com", "https://www.udemy.com"),
		html.Titulo("https://www.amazon.com.br", "https://www.youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
