package equacaobll

import (
	"POO/equacao"
	"POO/erro"
)

func ValidaDados(e *erro.Erro, eq equacao.Equacao) {
	e.SetErro(false)

	if eq.Delta() < 0 {
		e.SetErro(true)
		e.SetMens("Delta negativo! Não existem raízes reais para essa equação.")
	}
}
