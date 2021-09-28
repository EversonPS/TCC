package main

import "fmt"

func main() {
	// criação de variáveis
	var num1, num2 int

	// entrada de dados
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número (maior que o primeiro): ")
	fmt.Scanln(&num2)

	for num2 <= num1 { // enquanto o num2 for menor ou igual que o num1, repete-se as linhas de comando abaixo.
		//  em c# teríamos algo como while (num2 <= num1) ou algum laço de repetição semelhante.
		fmt.Print("Ele deve ser maior que o primeiro!!!\nDigite o segundo número: ")
		fmt.Scanln(&num2)
	}
}
