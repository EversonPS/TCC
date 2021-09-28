package main

import "fmt"

func main() {
	// criação de variáveis
	var resposta string = "S"
	var numero, fatorial, i int

	for resposta == "S" { // laço de repetição que faz com que, enquanto resposta = "S", o programa continuará rodando.

		// entrada de dados
		fmt.Print("\nDigite um número inteiro e positivo: ")
		fmt.Scanln(&numero)

		for numero < 0 { // enquanto o valor do número dado for menor que zero, o usuário terá de digitar novamente.
			fmt.Print("Erro!\nDigite um número inteiro e positivo: ")
			fmt.Scanln(&numero)
		}

		fatorial = numero
		//Saída de dados
		for i = numero - 1; i > 1; i-- { // este laço fará com que o fatorial do número será numero * (numero-1)... até numero * 1.
			fatorial = fatorial * i // este laço só ocorrerá caso o número seja inteiro, positivo e diferente de 0.
			// caso o número seja 1, será numero * numero, pois será numero * 1.
		}
		fmt.Printf("O fatorial de %d é %d", numero, fatorial)

		// Verificação se deseja continuar
		fmt.Print("\n\n Você deseja continuar? (Apenas S ou N serão aceitos como resposta!): ")
		fmt.Scanln(&resposta) //se a resposta for "S", o laço de repetição que inicia o programa é reiniciado.
		// se for "N", o programa acaba.

		for resposta != "S" && resposta != "N" { //se a resposta for diferente de "S" e "N", aparecerá uma mensagem de erro e
			fmt.Print("\n\n Erro!\nVocê deseja continuar?(Apenas S ou N serão aceitos como resposta!) ") // pedirá novamente.
			fmt.Scanln(&resposta)
		}
	}

}
