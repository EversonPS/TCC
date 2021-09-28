package main

import (
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

	// Variáveis dos compontentes do formulário
	var horista *walk.RadioButton
	var professor *walk.RadioButton
	var qtdH, vlrH *walk.NumberEdit
	var form *walk.MainWindow

	var calc float64

	// Janela principal
	MainWindow{
		// AssignTo apona para as váriaveis dos componentes do form
		AssignTo: &form,
		Title:    "Calculo de Salário Bruto",
		Size:     Size{400, 200},

		// Grid{Columns: 2} faz com que os itens se dividam no form 2 a 2
		Layout: Grid{Columns: 2},

		// Define o tipo da fonte e o seu tamanho
		Font: Font{Family: "Arial", PointSize: 14},

		// Define que todo o seu conteúdo estará dentro da janela principal
		Children: []Widget{

			// Grupo de Radio Button.
			RadioButtonGroupBox{

				// Título do grupo de radios buttons.
				Title:  "Você é:",
				Layout: Grid{Columns: 2},
				Buttons: []RadioButton{
					// Definindo dois radios buttons (Horista e Professor).
					{Text: "Horista", Value: "Hor", AssignTo: &horista},
					{Text: "Professor", Value: "Prof", AssignTo: &professor},
				},
			},

			Label{
				Text: "Digite a quantidade de horas/aulas:",
			},

			// Caixa de Texto numérico onde será digitado a quantidade de horas.
			NumberEdit{
				AssignTo: &qtdH,
			},

			Label{
				Text: "Digite o valor da hora/aula:",
			},

			// Caixa de Texto numérico onde será digitado o valor de horas trabalhadas.
			NumberEdit{
				AssignTo: &vlrH,

				// Define a quantidade de casas decimais e adiciona um texto no final da caixa.
				Suffix:   " BRL",
				Decimals: 2,
			},

			// Botão responsável por calcular o salário bruto
			PushButton{
				// Texto do botão
				Text: "Calcular",

				// Função quando o botão for clicado
				OnClicked: func() {

					// Checando se o radio button "professor" foi marcado
					if professor.Checked() == true {
						// Calculo do salario para professor
						calc = vlrH.Value() * qtdH.Value() * 1.25

						// Exibindo salário em uma caixa de mensagem
						walk.MsgBox(form, "Salário Bruto", "Salário Bruto: "+strconv.FormatFloat(calc, 'f', -1,
						64)+" Reais", walk.MsgBoxIconInformation)
					}

					// Verificando se o radio button "horista" foi marcado
					if horista.Checked() == true {
						// Calculo do salario para horista
						calc = vlrH.Value() * qtdH.Value()
						walk.MsgBox(form, "Salário Bruto", "Salário Bruto: "+strconv.FormatFloat(calc,
					    'f', -1, 64)+" Reais", walk.MsgBoxIconInformation)
					}

					/* Checando se nenhum dos radios buttons foram marcados,
					   se nao forem será exibido uma caixa de mensagem exibindo o erro */
					if horista.Checked() == false && professor.Checked() == false {
						walk.MsgBox(form, "Salário Bruto", "Selecione uma das opções para calcular o salário bruto!",
						walk.MsgBoxIconError)
					}
				},
			},

			// Botão Limpar, caso seja clicado limpará todos os campos do form
			PushButton{
				Text: "Limpar",
				OnClicked: func() {
					qtdH.SetValue(0)
					vlrH.SetValue(0)
					horista.SetChecked(false)
					professor.SetChecked(false)
				},
			},
		},
	}.Run()
}
