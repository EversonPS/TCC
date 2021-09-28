package livrobll

import (
	"POO/erro"
	"POO/livro"
)

func ValidaDados(e *erro.Erro, l livro.Livro) {
	e.SetErro(false)

	if l.GetCodigo() == "" {
		e.SetMens("O campo código é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetTitulo() == "" {
		e.SetMens("O campo titulo é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetAutor() == "" {
		e.SetMens("O campo autor é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetEditora() == "" {
		e.SetMens("O campo editora é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetAno() <= 0 {
		e.SetMens("O campo ano tem de ser maior do que 0")
		e.SetErro(true)
		return
	}
}

func ValidaCodigo(e *erro.Erro, l livro.Livro, cod string) {
	e.SetErro(false)
	if cod == "" {
		e.SetMens("O campo código é de preenchimento obrigatório")
		e.SetErro(true)
	} else {
		if l.GetCodigo() != cod {
			e.SetMens("Não há nenhum livro cadastrado com este código")
			e.SetErro(true)
		}
	}
}
