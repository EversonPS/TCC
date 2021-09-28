package triangulo

type Triangulo struct {
	altura float64
	base   float64
}

func (t *Triangulo) SetAltura(altura float64) {
	t.altura = altura
}

func (t Triangulo) GetAltura() float64 {
	return t.altura
}

func (t *Triangulo) SetBase(base float64) {
	t.base = base
}

func (t Triangulo) GetBase() float64 {
	return t.base
}

func (t Triangulo) GetArea() float64 {
	return t.base * t.altura / 2
}
