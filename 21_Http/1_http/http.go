package main

import (
	"log"
	"net/http"
)

// criando o controller da rota Home
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bem vindo ao nossa home page"))
}
func main() {
	//criando uma rota
	http.HandleFunc("/home", home)

	//criando um servidor (deve ser na ultima linha)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
