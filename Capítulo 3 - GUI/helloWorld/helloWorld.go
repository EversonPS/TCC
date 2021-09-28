package main

import (
	. "github.com/lxn/walk/declarative"
)

func main() {
	// Janela principal
	MainWindow{
		//titulo da janela
		Title: "Primeiro Programa",

		// Tamanho da janela
		Size: Size{300, 50},

		Layout: VBox{},
		Children: []Widget{
			// Label
			Label{
				// Texto que será exibido na label
				Text: "Hello World!",
			},
		},
	}.Run()
}
