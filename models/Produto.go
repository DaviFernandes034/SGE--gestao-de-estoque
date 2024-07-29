package models

import "time"


type Produtos struct{

	Id_produto int `json:"produto_id"`
	nome string `json:"nome"`
	preco float64 `json:"preco"`
	lote int `json:"lote"`
	validade time.Time `json:validade`
	categoria Categorias `json:"categoria"`
	

}