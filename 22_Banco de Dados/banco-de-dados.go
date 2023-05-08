package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //import implicito
)

func main() {
	//string de conexxao com o banco
	//user:pass@tcp(adressDb)/dataBaseName?charset=utf8mb4&parseTime=True&loc=Local"
	stringConexao := "root:@tcp(localhost:3306)/localdatabase?charset=utf8mb4&parseTime=True&loc=Local"

	//cria uma conexao com o banco
	db, erro := sql.Open("mysql", stringConexao)

	//tratando erros de driver
	if erro != nil {
		log.Fatal(erro)
	}
	//Após a execução da main, a conexxao é fechada
	defer db.Close()

	//Tratando erro de conexao com o banco
	if erro := db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("A conexão estar aberta")

	//Faz uma consulta no banco de dados
	linhas, erro := db.Query("SELECT * FROM usuarios")

	//Trata o erro caso a consulta apresente alguma falha
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(linhas)

	defer linhas.Close() //fecha as linha impedindo enumerações

}
