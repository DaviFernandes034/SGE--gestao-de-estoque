package dtos

import "time"

type ProdutosDTO struct {
	
	Status	   StatusDTO 		`json:"status"`
	Categoria  CategoriasDTO	`json:"categoria"`
	Nome       string    		`json:"nome"`
	Preco      float64   		`json:"preco"`
	Lote       string   		`json:"lote"`
	Validade   time.Time 		`json:"validade"`
}


func NewProdutoDTO(statusDto StatusDTO, categoriaDto CategoriasDTO, nome string, preco float64, lote string, validade time.Time) *ProdutosDTO {

	return&ProdutosDTO{
		 Status: statusDto,
		 Categoria: categoriaDto,
		 Nome: nome,
		 Preco: preco,
		 Lote: lote,
		 Validade: validade,
	}

}