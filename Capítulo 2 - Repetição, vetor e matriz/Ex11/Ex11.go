package main

import "fmt"

func main() {
	// declaração de variáveis
	var registro [20][15]string // declaração da matriz registro
	var nome string
	fileira := 0
	cadeira := 0
	a := false
	resp := "S"

	// entrada de dados (Nome, Fileira e Cadeira)
	fmt.Print("Nome: ")
	fmt.Scanln(&nome)

	for resp == "S" {
		a = false
		fmt.Print("\nFileira (1 a 20): ")
		fmt.Scanln(&fileira)
		for fileira <= 0 && fileira > 20 {
			fmt.Print("\nErro!\nFileira (1 a 20): ")
			fmt.Scanln(&fileira)
		}

		fmt.Print("Cadeira (1 a 15): ")
		fmt.Scanln(&cadeira)
		for cadeira <= 0 && cadeira > 15 {
			fmt.Print("\nErro!\nCadeira (1 a 15): ")
			fmt.Scanln(&cadeira)
		}

		fileira = fileira - 1
		cadeira = cadeira - 1

		// Verificando se a cadeira e fileira do cinema já está reservada
		for i := 0; i < 20; i++ {
			for j := 0; j < 15; j++ {
				// Se o local não estiver ocupado, será ser reservada com o bloco de instruções abaixo
				if registro[fileira][cadeira] == "" {
					registro[fileira][cadeira] = nome
					fmt.Print("\nA cadeira ", cadeira+1, " na fileira ", fileira+1, " foi reservada com sucesso")
					a = true
				}
			}
		}

		// Se já estiver ocupada, a seguinte mensagem será exibida
		if a == false {
			fmt.Print("\nReserva mal sucedida. Assento ocupado")
		}

		// Por fim, será perguntado se quer ou não reservar mais uma cadeira do cinema, se sim ("S"), se não ("N")
		fmt.Print("\nDeseja fazer uma nova reserva? (digite apenas S ou N): ")
		fmt.Scanln(&resp)
		for resp != "S" && resp != "N" {
			fmt.Print("\nErro!\nDeseja fazer uma nova reserva? (digite apenas S ou N): ")
			fmt.Scanln(&resp)
		}
	}
}
