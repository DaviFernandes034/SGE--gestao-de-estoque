package services

import (
	

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	
)


type CategoriaService struct {

	*BaseService[models.Categorias]
}


func NewCategoriaService(rc interfaces.RepositoryCrud[models.Categorias]) interfaces.ServiceCrud[models.Categorias] {

	return &CategoriaService{
		BaseService: NewBaseService[models.Categorias](rc),
	}
}

