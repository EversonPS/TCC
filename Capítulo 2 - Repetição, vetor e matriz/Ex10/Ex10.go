package main

import "fmt"

func main() {
	// Declaração de Variáveis
	var nome [10]string
	var idade [10]string
	var nomePesq string
	resp := "S"

	// Entrada de dados
	for i := 0; i <= 9; i++ {
		fmt.Print("\nDigite o nome da ", i+1, "° pessoa: ")
		fmt.Scanln(&nome[i])
		fmt.Print("Digite a idade da ", i+1, "° pessoa: ")
		fmt.Scanln(&idade[i])
	}

	for resp == "S" {
		a := false
		fmt.Print("\nDigite o nome da pessoa que deseja saber a idade: ")
		fmt.Scanln(&nomePesq)

		// Saída de dados
		for i := 0; i <= 9; i++ {
			// Comparação para ver se o nome pesquisado foi um dos 10 nomes digitados anteriormente
			if nome[i] == nomePesq {
				fmt.Println("A idade do ", nomePesq, "é de: ", idade[i], "anos")
				a = true
			}
		}

		/* se a for falso, significa que não entrou na estrutura de condição,
		portanto o nome não foi cadastrado */
		if a == false {
			fmt.Println("Pessoa não localizada!!")
		}

		// Verificação para uma nova consulta
		fmt.Print("\nDeseja fazer uma nova consulta? (digite apenas S ou N): ")
		fmt.Scanln(&resp)
		for resp != "S" && resp != "N" {
			fmt.Print("\nErro!\nDeseja fazer uma nova consulta? (digite apenas S ou N): ")
			fmt.Scanln(&resp)
		}
	}
}
