package main

import "fmt"

func main() {
	// criação de variáveis
	var sexo string

	// entrada de dados
	fmt.Print("Digite o seu sexo: (F ou M)")
	fmt.Scanln(&sexo)

	for sexo != "f" && sexo != "F" && sexo != "m" && sexo != "M" { // enquanto valor de sexo for diferente de "f", "F", "m" e "M", o laço
		// se repetirá.
		fmt.Print("Sexo inválido!\nDigite o seu sexo: ")
		fmt.Scanln(&sexo)
	}

}
