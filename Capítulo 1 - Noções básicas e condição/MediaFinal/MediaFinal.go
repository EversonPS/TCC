package main

import "fmt"

func main() {
	// declaração de variaveis
	var n1, n2, n3, n4 float64

	// Entrada de dados do usuário
	fmt.Print("Nota do 1° bimestre: ")
	fmt.Scanln(&n1)
	fmt.Print("Nota do 2° bimestre: ")
	fmt.Scanln(&n2)
	fmt.Print("Nota do 3° bimestre: ")
	fmt.Scanln(&n3)
	fmt.Print("Nota do 4° bimestre: ")
	fmt.Scanln(&n4)

	//calculo e exibição da média aritmética
	media := (n1 + n2 + n3 + n4) / 4
	fmt.Printf("Média final: %.1f. ", media)

	// estrutura condicional com if/else para exibir a situação do aluno
	if media < 3 {
		fmt.Print("Reprovado!\n")
	} else if media >= 3 && media <= 5 { // && é o operador lógico and
		fmt.Print("Recuperação!\n")
	} else {
		fmt.Print("Aprovado!\n")
	}
}
