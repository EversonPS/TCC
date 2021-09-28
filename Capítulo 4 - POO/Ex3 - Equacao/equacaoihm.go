package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"POO/equacao"
	"POO/equacaobll"
	"POO/erro"
)

func main() {
	var vlrA, vlrB, vlrC, x1, x2 *walk.NumberEdit
	var form *walk.MainWindow

	MainWindow{
		AssignTo: &form,
		Title:    "Equação do segundo grau",
		Size:     Size{400, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Valor de a:",
			},

			NumberEdit{
				AssignTo: &vlrA,
			},

			Label{
				Text: "Valor de b:",
			},

			NumberEdit{
				AssignTo: &vlrB,
			},

			Label{
				Text: "Valor de c:",
			},

			NumberEdit{
				AssignTo: &vlrC,
			},

			PushButton{
				Text: "Calcular",
				OnClicked: func() {
					// Instanciando structs
					erro := new(erro.Erro)
					equacao := new(equacao.Equacao)

					// Setando valores nos campos da struct Equacao
					equacao.SetA(int(vlrA.Value()))
					equacao.SetB(int(vlrB.Value()))
					equacao.SetC(int(vlrC.Value()))

					// Validando dados inseridos pelo usuário
					equacaobll.ValidaDados(erro, *equacao)

					/* Caso os dados inseridos estejam corretos será exibido o valor de x1 e x2,
					caso contrário ele mostrará uma mensagem de erro que foi setado na struct Erro
					pela método ValidaDados presente no package equacaobll */
					if erro.GetErro() {
						walk.MsgBox(form, "Erro", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {
						x1.SetValue(equacao.X1())
						x2.SetValue(equacao.X2())
						vlrA.SetEnabled(false)
						vlrB.SetEnabled(false)
						vlrC.SetEnabled(false)
					}
				},
			},

			PushButton{
				Text: "Limpar",
				OnClicked: func() {
					vlrA.SetValue(0)
					vlrB.SetValue(0)
					vlrC.SetValue(0)
					x1.SetValue(0)
					x2.SetValue(0)
					vlrA.SetEnabled(true)
					vlrB.SetEnabled(true)
					vlrC.SetEnabled(true)
				},
			},

			Label{
				Text: "x1: ",
			},
			NumberEdit{
				AssignTo: &x1,
				Enabled:  false,
				Decimals: 2,
			},

			Label{
				Text: "x2: ",
			},
			NumberEdit{
				AssignTo: &x2,
				Enabled:  false,
				Decimals: 2,
			},
		},
	}.Run()
}
