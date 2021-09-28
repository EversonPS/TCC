package main

import (
	"fmt"
)

func main() {
	//declaração de variáveis
	var numero [10]int //criação de uma matriz de tamanho 10. em c# seria "int[] numero = new float[10];"
	maior := 0
	menor := 0
	//entrada de dados
	for i := 0; i <= 9; i++ { // laço de repetição com i começando em 0, pois a primeira posição de uma matriz começa no índice 0 ([0])
		//o laço termina com i = 10, pois a condição é que o laço continue enquanto i for menor ou igual a 9.
		fmt.Print("Digite o ", i+1, "° número: ")
		fmt.Scanln(&numero[i])

		if numero[i] > maior {
			maior = numero[i]
			if i == 0 {
				menor = maior
			}
		}
		if numero[i] < menor {
			menor = numero[i]
		}
	}
	//saída de dados
	fmt.Print("\n\nO menor número digitado é: ", menor)
	fmt.Print("\nO maior número digitado é: ", maior, "\n")

}
