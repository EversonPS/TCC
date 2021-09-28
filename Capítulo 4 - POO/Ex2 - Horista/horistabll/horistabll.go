package horistabll

import (
	"POO/erro"
	"POO/horista"
)

func ValidaDados(e *erro.Erro, h horista.Horista) {
	e.SetErro(false)

	if h.GetQtd() <= 0 && h.GetValor() <= 0 {
		e.SetErro(true)
		e.SetMens("O campo QUANTIDADE DE HORAS e VALOR DA HORA devem ser maiores que 0")
	} else {
		if h.GetQtd() <= 0 {
			e.SetErro(true)
			e.SetMens("O campo QUANTIDADE DE HORAS deve ser maior que 0")
		} else if h.GetValor() <= 0 {
			e.SetErro(true)
			e.SetMens("O campo VALOR DA HORA deve ser maior que 0")
		}
	}
}
