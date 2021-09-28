package horista

type Horista struct {
	valor float64
	qtd   int
}

func (h *Horista) SetQtd(qtd int) {
	h.qtd = qtd
}

func (h Horista) GetQtd() int {
	return h.qtd
}

func (h *Horista) SetValor(vlr float64) {
	h.valor = vlr
}

func (h Horista) GetValor() float64 {
	return h.valor
}

func (h Horista) SalarioBruto() float64 {
	return float64(h.qtd) * h.valor / 2
}
