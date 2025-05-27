package dtos



type CategoriasDTO struct{

	Nome string `json:"nome"`

}

func NewCategoriaDto(nome string) *CategoriasDTO{

	return &CategoriasDTO{
		Nome: nome,
	}
}
