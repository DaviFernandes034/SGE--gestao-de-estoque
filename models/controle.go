package models

import "time"


type Controle struct{

	Id_controle int `json:"controle_id"`
	Id_produto Produtos  `json:"produto"`
	Id_status Status `json:"status"`
	quantidade int `json:"quantidade"`
	data time.Time `json:"data"`
}