package models


type Categorias struct{

	Id_categoria int64 `json:"categoria_id`
	Nome string `json:"nome"`
}


type CategoriasRequest struct{

	Nome string `json:"nome"`

}