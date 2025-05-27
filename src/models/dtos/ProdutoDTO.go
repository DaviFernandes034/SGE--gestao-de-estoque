package dtos

import "time"

type ProdutosDTO struct {
	
	Status	   StatusDTO 		`json:"status`
	Categoria  CategoriasDTO	`json:"categoria"`
	Nome       string    		`json:"nome"`
	Preco      float64   		`json:"preco"`
	Lote       string   		`json:"lote"`
	Validade   time.Time 		`json:"validade"`
}