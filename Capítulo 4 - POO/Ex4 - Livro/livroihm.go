package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"POO/erro"
	"POO/livro"
	"POO/livrobll"
)

func main() {
	var cod, titulo, editora, autor *walk.TextEdit
	var consultar *walk.PushButton
	var ano *walk.NumberEdit
	var form *walk.MainWindow

	// Instanciando struct Livro e Erro. "livroaux" será utilizado apenas para validação
	livroAux := new(livro.Livro)
	erroCadastra := new(erro.Erro)

	// Instanciando struct Livro. Os valores de "livro" serão utilizados pós validação
	livro := new(livro.Livro)
	MainWindow{
		AssignTo: &form,
		Title:    "Cadastro de Livro",
		Size:     Size{300, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Código:",
			},

			TextEdit{
				Text:     "",
				AssignTo: &cod,
			},

			Label{
				Text: "Titulo:",
			},

			TextEdit{
				Text:     "",
				AssignTo: &titulo,
			},

			Label{
				Text: "Autor",
			},

			TextEdit{
				Text:     "",
				AssignTo: &autor,
			},

			Label{
				Text: "Editora",
			},

			TextEdit{
				Text:     "",
				AssignTo: &editora,
			},

			Label{
				Text: "Ano",
			},

			NumberEdit{
				AssignTo:           &ano,
				RightToLeftReading: false,
			},

			PushButton{
				ColumnSpan: 2,
				Text:       "Cadastrar",
				OnClicked: func() {
					// Setando valores no "livroaux"
					livroAux.SetTitulo(titulo.Text())
					livroAux.SetAutor(autor.Text())
					livroAux.SetCodigo(cod.Text())
					livroAux.SetEditora(editora.Text())
					livroAux.SetAno(int(ano.Value()))

					// Verificando se o que foi digitado está correto ou não
					livrobll.ValidaDados(erroCadastra, *livroAux)

					/* Caso esteja errado, será exibido uma mensagem de erro.
					Caso contrário além de informar que os dados estão corretos, os dados de "livroaux" serão setados
					na instância "livro" */
					if erroCadastra.GetErro() {
						walk.MsgBox(form, "Erro", erroCadastra.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {
						consultar.SetEnabled(true)
						walk.MsgBox(form, "Boa!", "Os dados inseridos com sucesso!", walk.MsgBoxStyle(walk.MsgBoxIconInformation))
						livro.SetTitulo(livroAux.GetTitulo())
						livro.SetAutor(livroAux.GetAutor())
						livro.SetCodigo(livroAux.GetCodigo())
						livro.SetEditora(livroAux.GetEditora())
						livro.SetAno(livroAux.GetAno())
					}
				},
			},

			PushButton{
				AssignTo:   &consultar,
				ColumnSpan: 2,
				Text:       "Consultar",
				Enabled:    false,
				OnClicked: func() {
					// Criando nova instância da struct Erro
					erro := new(erro.Erro)

					// Verificará se o código digitado na caixa de texto é igual ao código do livro cadastrado.
					livrobll.ValidaCodigo(erro, *livro, cod.Text())

					/* Se der tudo certo será mostrado os dados do livro cadastrado, senão será mostrado uma
					mensagem de erro */

					if erro.GetErro() {
						walk.MsgBox(form, "Erro", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {
						cod.SetText(livro.GetCodigo())
						titulo.SetText(livro.GetTitulo())
						autor.SetText(livro.GetAutor())
						editora.SetText(livro.GetEditora())
						ano.SetValue(float64(livro.GetAno()))
					}
				},
			},

			PushButton{
				ColumnSpan: 2,
				Text:       "Limpar",
				OnClicked: func() {
					cod.SetText("")
					titulo.SetText("")
					autor.SetText("")
					editora.SetText("")
					ano.SetValue(0)
				},
			},
		},
	}.Run()
}
