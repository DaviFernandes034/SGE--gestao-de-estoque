package configs

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

type InitConnection struct {
	Conn *Connection
}

func (init *InitConnection) Init() ( error){

	err:= godotenv.Load()
	if err != nil {

		return fmt.Errorf("erro ao carregar o arquivo .env: %w", err)
	}

	db, err:= conn()
	if err != nil {
		return fmt.Errorf("erro ao chamar funcao Conn: %w", err)
	}

	log.Println("INFO: função para iniciar o banco de dados e carregar os arquivos .env")
	
	init.Conn = db

	return nil

}