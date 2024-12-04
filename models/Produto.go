package models

import (
	"time"

	
)

type Produtos struct{

	Id_produto int `json:"produto_id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
	Lote string `json:"lote"`
    Validade time.Time `json:validade`
	CategoriaId int64 `json:categoriaId`
	

}


type ProdutosRequest struct{

	Id_produto int `json:"produto_id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
	Lote string `json:"lote"`
    Validade time.Time `json:validade` 
	Categoria Categorias `json:"categoria"`
	
	

}