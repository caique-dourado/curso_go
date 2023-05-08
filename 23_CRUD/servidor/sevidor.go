package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// 1- Lendo a requisição recebida
	corpoRequisicao, erro := io.ReadAll(r.Body)

	//2- criando uma instacia vazia para armazenar os dados da requisiçao lida
	var usuario usuario
	if erro != nil {
		w.Write([]byte("Erro ao obter corpo da requisição"))
		return
	}

	// 3- convertendo os dados da requisição lidos (JSON) em um struct
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario em um struct"))
		return
	}

	// 4- abrindo a conexão com o banco de dados
	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	//5- protegendo as requisições do banco de sql injector usando o Prepare()
	statement, erro := db.Prepare("INSERT INTO usuarios(nome, email) VALUES (?,?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	//6- Executando o statement com os dados do struct a ser enviados para o banco
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}

	//7- passando o status code para o cabeçalho da requisição
	w.WriteHeader(http.StatusCreated)

	//8- msg de confirmação
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao recuperar o id inserido"))
		return
	}
	w.Write([]byte(fmt.Sprintf("dado inserido: %d", idInserido)))

}
