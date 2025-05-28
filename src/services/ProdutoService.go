package services

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type ProdutoService struct {


	*BaseService[models.Produtos]
}

func NewProdutoService(rp interfaces.RepositoryCrud[models.Produtos]) interfaces.ServiceCrud[models.Produtos ]{

	return &ProdutoService{

		BaseService: NewBaseService[models.Produtos](rp),
	}
}
