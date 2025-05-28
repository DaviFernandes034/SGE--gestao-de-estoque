package configs

import (
	
	"fmt"
	"log"
)

//struct da conexao para iniciar o banco de dados
type InitConnection struct {
	Conn *Connection
}

func (init *InitConnection) Init() ( *Connection, error){


	db, err:= conn()//função conn vindo do arquivo que carrega a inicialização do banco de dados
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar funcao Conn: %w", err)
	}

	log.Println("INFO: função para iniciar o banco de dados")
	
	init.Conn = db //adicionando a conexão da struct "Connection" a struct InitConnection

	return db, nil

}