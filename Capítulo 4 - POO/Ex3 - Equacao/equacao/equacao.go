package equacao

import "math"

type Equacao struct {
	vlrA int
	vlrB int
	vlrC int
}

func (e *Equacao) SetA(a int) {
	e.vlrA = a
}

func (e Equacao) GetA() int {
	return e.vlrA
}

func (e *Equacao) SetB(b int) {
	e.vlrB = b
}

func (e Equacao) GetB() int {
	return e.vlrB
}

func (e *Equacao) SetC(c int) {
	e.vlrC = c
}

func (e Equacao) GetC() int {
	return e.vlrC
}

func (e Equacao) Delta() float64 {
	a := e.GetA()
	b := e.GetB()
	c := e.GetC()
	return float64((b * b) - 4*a*c)
}

func (e Equacao) X1() float64 {
	a := e.GetA()
	b := e.GetB()
	return (-float64(b) + (math.Sqrt(e.Delta()))) / (2 * float64(a))
}

func (e Equacao) X2() float64 {
	a := e.GetA()
	b := e.GetB()
	return (-float64(b) - (math.Sqrt(e.Delta()))) / (2 * float64(a))
}
