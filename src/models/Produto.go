package models

import (
	"time"

	"github.com/google/uuid"
)

type Produtos struct{

	Id_produto uuid.UUID `json:"produto_id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
	Lote string `json:"lote"`
    Validade time.Time `json:"validade"`
	CategoriaId int64 `json:"categoriaId"`
	StatusId int64 `json:"statusId"`
	

}


