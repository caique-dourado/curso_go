package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// A função externa Conectar retorna uma conexao(*sql.DB) ou um erro (error)
func Conectar() (*sql.DB, error) {
	//string de conexão
	uri := "root:@tcp(localhost:3306)/localdatabase?charset=utf8mb4&parseTime=True&loc=Local"

	//abrindo uma conexão
	db, erro := sql.Open("mysql", uri)

	//verificando erros externos
	if erro != nil {
		//retorna o valor 0 para o ponteiro *sql.DB e um erro para o error
		return nil, erro
	}

	//validando a conexão
	if erro = db.Ping(); erro != nil {
		///retorna o valor 0 para o ponteiro *sql.DB e um erro de conexão para o error
		return nil, erro

	}

	//SE TUDO DER CERTO ↓
	//retorna o banco para o ponteiro *sql.DB e  o valor 0(em branco) para o error
	return db, nil
}
