package main

import "fmt"

func main() {

	// Criando váriaveis do tipo float para números com possíveis casas decimais
	var altura float32
	var base float32
	var area float32

	// fmt.Print responsável pela saída de dados, (Console.Write em C#)
	fmt.Print("Base do triângulo: ")
	fmt.Scanln(&base) // fmt.Scanln é responsavel pela entrada de dados do usuario, (Console.ReadLine em C#)
	fmt.Print("Altura do triângulo: ")
	fmt.Scanln(&altura)

	area = base * altura / 2

	/* printf é usado para formatar o texto. Com "%.1f"	mostrando o lugar
	onde a variavel irá ser exibida, o seu tipo (float) e com uma casa decimal. */
	fmt.Printf("Area: %.1f", area) // % funciona de maneira muito parecida com o placeholder {0} em c#
	fmt.Println("")
}
