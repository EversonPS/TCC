package main

import (
	"fmt"
	"math"
)

func main() {
	// definição de variavel inteira (int)
	var grau int

	var rad float64

	fmt.Print("Angulo em °: ")
	fmt.Scanln(&grau)

	/* math.PI, math.sin e math.cos são funções matemáticas
	que retornam repectivamente o PI(3.14), seno e cosseno. */
	rad = float64(grau) * math.Pi / 180
	cosseno := math.Cos(float64(rad))
	seno := math.Sin(float64(rad))

	fmt.Printf("cos: %.2f\nsen: %.2f\n", cosseno, seno)
}
