package main

import (
	"fmt"
	"math" // importamos a biblioteca math para podermos usar funções dela. neste programa é usada a função "Math.Pow(x,y)", que retorna x**y (x elevado a y)
) //  encontramos a mesma função em C#, da mesma maneira.

func main() {
	// criação de variáveis
	var n, i int
	var x float64 = 1

	// entrada de dados
	fmt.Print("Digite os n primeiros termos da série: ") //"fmt.Print" (GO) = "Console.Write" (C#)
	fmt.Scan(&n)                                         // "fmt.Scanln", "fmt.Scan" e "fmt.Scanf" se assemelham ao "Console.ReadLine()"

	// saida de dados
	for i = 1; i <= n; i++ { // laço de repetição, onde a variável i começa com valor 1 e vai até o valor de n+1.
		// assim que a variável chegar ao valor n + 1 o laço de repetição termina.
		if i != n {
			fmt.Printf("%0.f, ", math.Pow(x, 2)+1) // aqui, usamos o Printf, onde tudo é formatado para string.
		} else { //  "%0.f" se refere à variável float que é retornada de "math.Pow(x, 2)+1".
			fmt.Printf("%0.f", math.Pow(x, 2)+1) //   "0." é uma limitação de casas depois da vírgula. no caso, não teremos nenhum número após a vírgula.
		} //    podemos colocar qualquer número para limitar as casas depois da vírgula.
		x++ // x = x + 1. da mesma forma é feito em C#.

	}
}
