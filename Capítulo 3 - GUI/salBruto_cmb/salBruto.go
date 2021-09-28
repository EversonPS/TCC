package main

import (
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var qtdH, vlrH *walk.NumberEdit
	var form *walk.MainWindow
	var combo *walk.ComboBox
	var calc float64

	MainWindow{
		AssignTo: &form,
		Title:    "Calculo de Salário Bruto",
		Size:     Size{400, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Insira: ",
			},
			ComboBox{
				AssignTo: &combo,

				// Não deixará o usuário digitar o que quiser
				Editable: false,

				// o que aparecerá dentro da combobox
				Model:    []string{"Horista", "Professor"},
			},

			Label{
				Text: "Digite a quantidade de horas/aulas:",
			},
			NumberEdit{
				AssignTo: &qtdH,
			},

			Label{
				Text: "Digite o valor da hora/aula:",
			},
			NumberEdit{
				AssignTo: &vlrH,
				Suffix:   " BRL",
				Decimals: 2,
			},

			PushButton{
				Text: "Calcular",
				OnClicked: func() {
					// Checando qual opção foi marcada no checkbox
					if combo.Text() == "Horista" {
						calc = vlrH.Value() * qtdH.Value()
						walk.MsgBox(form, "Salário Bruto", "Salário Bruto: "+strconv.FormatFloat(calc,
							'f', -1, 64)+" Reais", walk.MsgBoxIconInformation)
					}

					if combo.Text() == "Professor" {
						calc = vlrH.Value() * qtdH.Value() * 1.25
						walk.MsgBox(form, "Salário Bruto", "Salário Bruto: "+strconv.FormatFloat(calc,
							'f', -1, 64)+" Reais", walk.MsgBoxIconInformation)
					}

					if combo.Text() != "Professor" && combo.Text() != "Horista" {
						walk.MsgBox(form, "Salário Bruto", "Selecione uma das opções para calcular o salário bruto!",
							walk.MsgBoxIconError)
					}
				},
			},

			PushButton{
				Text: "Limpar",
				OnClicked: func() {
					qtdH.SetValue(0)
					vlrH.SetValue(0)

					// Limpa a comboBox
					combo.SetCurrentIndex(-1)
				},
			},
		},
	}.Run()
}
