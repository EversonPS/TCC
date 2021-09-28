package main

import (
	"POO/erro"
	"POO/horista"
	"POO/horistabll"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var qtdH, vlrH, Salario *walk.NumberEdit
	var form *walk.MainWindow

	MainWindow{
		AssignTo: &form,
		Title:    "Calculo de Salário Bruto",
		Size:     Size{400, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Digite a quantidade de horas:",
			},
			NumberEdit{
				AssignTo: &qtdH,
			},

			Label{
				Text: "Digite o valor da hora:",
			},
			NumberEdit{
				AssignTo: &vlrH,
				Suffix:   " BRL",
				Decimals: 2,
			},

			PushButton{
				Text: "Calcular",
				OnClicked: func() {

					// Criando instâncias das structs Erro e Horista
					erro := new(erro.Erro)
					horista := new(horista.Horista)

					// Setando o que os valores digitados pelo usuário nos campos da struct Horista
					horista.SetQtd(int(qtdH.Value()))
					horista.SetValor(vlrH.Value())

					// Validando se o que o usuário digitou está correto
					horistabll.ValidaDados(erro, *horista)

					// Caso esteja errado será exibido uma mensagem de erro, caso contrário será exibido o salário na tela
					if erro.GetErro() {
						walk.MsgBox(form, "Erro!", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {
						Salario.SetValue(horista.SalarioBruto())
						qtdH.SetEnabled(false)
						vlrH.SetEnabled(false)
					}
				},
			},

			PushButton{
				Text: "Limpar",
				OnClicked: func() {
					qtdH.SetValue(0)
					vlrH.SetValue(0)
					Salario.SetValue(0)
					vlrH.SetEnabled(true)
					qtdH.SetEnabled(true)
				},
			},

			Label{
				Text: "Salário Bruto: ",
			},
			NumberEdit{
				AssignTo: &Salario,
				Suffix:   " BRL",
				Decimals: 2,
				Enabled:  false,
			},
		},
	}.Run()
}
