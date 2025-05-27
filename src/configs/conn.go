package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

type Connection struct {

	Db *sql.DB

}


func  conn() (*Connection, error) {
	// Obter variáveis de ambiente
	db_server:= os.Getenv("DB_SERVER")
	db_port:= os.Getenv("DB_PORT")
	db_database:= os.Getenv("DATABASE")


	//STRING CONEXÃO
	stringConn:= fmt.Sprintf("server=%s;port=%s;database=%s;trusted_connection=yes", db_server,db_port,
	db_database)

	db, err:= sql.Open("sqlserver", stringConn)
	if err != nil {

		return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %w", err)
	}

	err = db.Ping()
	if err != nil {

		return nil, fmt.Errorf("erro ao verificar se a conexão ainda está ativa ")
	}

	log.Println("INFO: conexão ao banco de dados funcionando!!")

	return &Connection{
		Db: db,
	}, nil

}
