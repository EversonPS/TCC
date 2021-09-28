package main

import (
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var cons *walk.NumberEdit
	var form *walk.MainWindow
	var list *walk.TextEdit
	var calcular *walk.PushButton

	var calc, i float64

	MainWindow{
		AssignTo: &form,
		Title:    "Calculo Tabuada",
		Size:     Size{250, 400},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Insira um valor: ",
			},

			// Campo que armazenará o valor da tabuda
			NumberEdit{
				AssignTo: &cons,
			},

			// Botão Calcular
			PushButton{
				AssignTo:   &calcular,
				ColumnSpan: 2,
				Text:       "Calcular",
				OnClicked: func() {
					// desabilitando o botão "calcular" até que o usuário clique no botão "limpar"
					calcular.SetEnabled(false)

					// Código gerador da tabuada
					for i = 1; i <= 10; i++ {
						// Calculo do valor recebido pelo usuario multiplicado pelas iterações do for
						calc = cons.Value() * i

						// Setando o texto na TextEdit onde será exibido a tabuada e pulando uma linha à cada iteração
						list.SetText(list.Text() + strconv.FormatFloat(cons.Value(), 'f', -1,
							64) + " * " + strconv.FormatFloat(i, 'f', -1, 64) + " = " + strconv.FormatFloat(calc,
							'f', -1, 64) + "\r\n")
					}
				},
			},

			// Text edit onde será exibida a tabuada
			TextEdit{
				ColumnSpan: 2,
				AssignTo:   &list,
				Text:       "",

				// Código que deixa a text edit apenas para leitura impossibilitando que o usuario insira texto
				ReadOnly: true,
			},

			// Limpando os campos do form
			PushButton{
				ColumnSpan: 2,
				Text:       "Limpar",
				OnClicked: func() {
					list.SetText("")
					cons.SetValue(0)

					// Após a limpeza habilitando o botão "calcular"
					calcular.SetEnabled(true)
				},
			},
		},
	}.Run()
}
