package main

import "fmt"

func main() {
	//declaração de variáveis
	var numero [10]float64 //criação de uma matriz de tamanho 10. em c# seria "int[] numero = new float[10];"
	var soma float64 = 0
	//entrada de dados
	for i := 0; i <= 9; i++ { // laço de repetição com i começando em 0, pois a primeira posição de uma matriz começa no índice 0 ([0])
		// laço termina com i = 10, pois a condição é que o laço continue enquanto i for menor ou igual a 9.
		fmt.Print("Digite o ", i+1, "° número: ")
		fmt.Scanln(&numero[i])
	}

	//saída de dados
	for i := 9; i >= 0; i-- { // laço de repetição com i começando em 9, logo, o índice 9 é o valor da posição 10 (décimo número armazenado)
		soma += numero[i] // então somaremos à variável soma (que vale 0) os valores de cada posição da matriz numero.
	}
	media := soma / 10
	fmt.Printf("\nA média dos números inseridos é %.2f!!\n", media)

}
