package repository

import (
	"log"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
)

type Repository struct {
	Db *configs.Connection
}

func NewRepository(db *configs.Connection)*Repository{

	log.Println("INFO: repository conectado com o banco de dados")
	return &Repository{
		Db: db,
	}


}