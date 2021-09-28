package main

import "fmt"

func main() {
	// criação de variáveis
	var num1, num2, aux int
	num1 = 0
	num2 = 1

	// laço de repetição, onde enquanto num2 for menor ou igual a 1000, continuará ocorrendo.
	for num2 <= 1000 {
		aux = num1
		num1 = num2
		num2 = num1 + aux
		if num2 < 1000 { // para assegurarmos que só teremos os números menores que 1000, criamos este if.
			fmt.Println(num2)
		}
	}
}
