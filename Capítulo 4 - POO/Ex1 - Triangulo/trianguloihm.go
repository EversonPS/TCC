package main

import (
	"POO/erro"
	"POO/triangulo"
	"POO/triangulobll"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var altura *walk.NumberEdit
	var base *walk.NumberEdit
	var resultado *walk.NumberEdit
	var calcular *walk.PushButton
	var form *walk.MainWindow

	MainWindow{
		AssignTo: &form,
		Title:    "Calculo Área",
		Size:     Size{250, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Base: ",
			},

			NumberEdit{
				AssignTo: &base,
			},

			Label{
				Text: "Altura: ",
			},

			NumberEdit{
				AssignTo: &altura,
			},

			PushButton{
				AssignTo: &calcular,
				Text:     "Calcular",
				OnClicked: func() {
					triangulo := new(triangulo.Triangulo)
					erro := new(erro.Erro)

					// setando valores da base e altura na struct triangulo
					triangulo.SetBase(base.Value())
					triangulo.SetAltura(altura.Value())

					// validando se o que o usuário digitou está correto
					triangulobll.ValidaDados(erro, *triangulo)

					/* se o que foi digitado esteja errado, será exibido uma mensagem de erro, caso
					contrário exibirá a área do triângulo */
					if erro.GetErro() {
						walk.MsgBox(form, "Erro!", erro.GetMens(), walk.MsgBoxIconError)
					} else {
						resultado.SetValue(triangulo.GetArea())
						base.SetEnabled(false)
						altura.SetEnabled(false)
					}
				},
			},

			PushButton{
				Text: "Limpar",
				OnClicked: func() {
					altura.SetValue(0)
					base.SetValue(0)
					resultado.SetValue(0)
					base.SetEnabled(true)
					altura.SetEnabled(true)
				},
			},

			Label{
				Text: "Área: ",
			},

			NumberEdit{
				AssignTo: &resultado,
				ReadOnly: true,
				Enabled:  false,
			},
		},
	}.Run()
}
