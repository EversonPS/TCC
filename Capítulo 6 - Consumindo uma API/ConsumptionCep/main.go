package main

// importação de bibliotecas
import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// estrutura JSON que recebemos através da API
type ConsultaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// setando o valor das variáveis
	var form *walk.MainWindow
	var cep, uf, cidade, bairro, rua *walk.TextEdit
	var consultar *walk.PushButton
	MainWindow{
		AssignTo: &form,
		Title:    "Consultar endereço",
		Size:     Size{500, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{

			Label{
				Text: "Digite seu CEP:",
			},

			// Campo CEP
			TextEdit{
				AssignTo: &cep,
				Text:     "",
			},

			// Botão de consulta
			PushButton{
				AssignTo:   &consultar,
				ColumnSpan: 2,
				Text:       "Consultar Endereço",
				OnClicked: func() {
					cepDigitado := cep.Text()
					url := "https://viacep.com.br/ws/" + cepDigitado + "/json/"

					// O método http.Get nos traz os valores que a API nos manda, caso dê erro, será retornado na segunda variável.
					response, _ := http.Get(url)

					// O método ioutil.ReadAll serve para convertermos o response.Body (corpo do response), que vem como formato de bytes, para string.
					responseData, _ := ioutil.ReadAll(response.Body)

					// Iniciamos uma variável do tipo ConsultaCEP
					var consulta ConsultaCEP

					// A função json.Unmarshal transforma nossa string para os valores recebidos em cada campo, indo diretamente para cada campo setado pelo `json:"nome do campo"`
					json.Unmarshal(responseData, &consulta)

					// Se os valores retornados forem vazios significa que o CEP inserido é inválido.
					if consulta.Uf == "" && consulta.Localidade == "" && consulta.Bairro == "" && consulta.Logradouro == "" {
						walk.MsgBox(form, "Erro!", "Erro ao consultar os dados!", walk.MsgBoxStyle(walk.MsgBoxIconError))
					}

					// Setando os valores recebidos para os campos visuais de texto.
					uf.SetText(consulta.Uf)
					cidade.SetText(consulta.Localidade)
					bairro.SetText(consulta.Bairro)
					rua.SetText(consulta.Logradouro)
				},
			},

			Label{
				Text: "Uf:",
			},

			TextEdit{
				AssignTo: &uf,
				Text:     "",
				Enabled:  false,
			},

			Label{
				Text: "Cidade:",
			},

			TextEdit{
				AssignTo: &cidade,
				Text:     "",
				Enabled:  false,
			},

			Label{
				Text: "Bairro:",
			},

			TextEdit{
				AssignTo: &bairro,
				Text:     "",
				Enabled:  false,
			},

			Label{
				Text: "Rua:",
			},

			TextEdit{
				AssignTo: &rua,
				Text:     "",
				Enabled:  false,
			},
		},
	}.Run()
}
