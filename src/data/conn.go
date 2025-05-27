package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

func Conn() (*sql.DB, error) {
	// Obter variáveis de ambiente
	server := os.Getenv("DB_SERVER")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	// Configuração da string de conexão
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, server, port, database)

	// Abrindo a conexão
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("erro na conexão ao banco de dados: ", err)
	}

	

	// Verificando se a conexão está ativa
	err = db.Ping()
	if err != nil {
		log.Fatal("erro ao verificar a conexao com o banco de dados: ", err)
	}

	fmt.Println("conexão estabelecida")

	return db, nil
}
