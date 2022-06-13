package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

/*
  template.Must: Encapsula todos os templates e de volve dois retornos ( o template e um erro)
  template.ParseGlob : recebe o caminho dos templates e carrega dos os arquivos com extensão HTML
  varialve temp recebe todas as paginas html
*/
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	/*
	  Atende e encaminha a requisição para a paigina index, recebenco como parametro o primeiro caminho
	  para atendimento que é a "/" mandando essa requisição para o index
	*/
	http.HandleFunc("/", index)
	//ouvir e responder a requisição: subindo um servidor.
	http.ListenAndServe(":8000", nil)

}
func index(w http.ResponseWriter, r *http.Request) {
	produto := []Produto{{"Máquina de tatoo", "Máquina rotativa", 3500, 10},
		{"Máquina de tatoo", "Máquina Bobina", 1500, 5},
		{"Máquina de tatoo", "Pen", 8000, 3},
		{"Fonte de tatoo", "Fonte Analógica 2 canais", 250, 8}}
	fmt.Println(produto)
	temp.ExecuteTemplate(w, "Index", produto)
}
