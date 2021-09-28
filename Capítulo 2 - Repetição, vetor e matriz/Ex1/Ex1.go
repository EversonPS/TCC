package main

import "fmt"

func main() {
	// criação de variáveis
	var numero, i, multiplicação int

	// entrada de dados
	fmt.Print("Digite um número inteiro e positivo: ") //"fmt.Print" (GO) = "Console.Write" (C#)
	fmt.Scanln(&numero)                                // "fmt.Scanln", "fmt.Scan" e "fmt.Scanf" se assemelham ao "Console.ReadLine()"

	// saida de dados
	fmt.Println("Tabuada do número: ", numero)
	for i = 1; i <= 10; i++ { // laço de repetição, onde a variável i começa com valor 1 e vai até 10+1.
		multiplicação = numero * i                             // assim que a variável chegar ao valor 10+1 o laço de repetição termina.
		fmt.Printf("%d x %d = %d\n", numero, i, multiplicação) // o for em GO é praticamente idêntico ao do C#, onde a diferença está só
	} // na ausência dos parênteses após "for"
}
