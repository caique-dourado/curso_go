package main

import (
	"html/template"
	"log"
	"net/http"
)

//criando um struct para enviar os dados do go para o html

type usuario struct {
	Nome  string
	Email string
}

// template (deve ser criado dessa forma)
var templates *template.Template

// controller
func home(w http.ResponseWriter, r *http.Request) {
	//criando a instancia u de usuario
	u := usuario{
		"Caique Dourado",
		"caique@dourado.com.br",
	}
	//executando o template criado passando os dados da instancia
	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {
	//buscando todos os arquivos .html e adcionando-as como template
	templates = template.Must(template.ParseGlob("*.html"))

	//rota
	http.HandleFunc("/home", home)
	//servidor
	log.Fatal(http.ListenAndServe(":8000", nil))
}
