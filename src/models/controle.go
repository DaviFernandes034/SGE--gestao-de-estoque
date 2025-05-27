package models

import "time"


type Controle struct{

	Id_controle int64 `json:"Controle_id"`
	Id_produto int64  `json:"Produto_id"`
	Id_status int64 `json:"Status_id"`
	Quantidade int64 `json:"Quantidade"`
	Data time.Time `json:"Data"`
}


type ControleRequest struct{

	Id_controle int64 `json:"Controle_id"`
	Produto ProdutosRequest `json:Produto`
	Status Status	`json:Status`
	Quantidade int `json:"Quantidade"`
	Data time.Time `json:"Data"`

}