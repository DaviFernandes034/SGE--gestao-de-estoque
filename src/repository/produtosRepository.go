package repository

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
)

type ProdutoRepository struct {
	Repository
}

// Create implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) Create(entity models.Produtos) error {
	panic("unimplemented")
}

// Delete implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) Delete(entity models.Produtos) error {
	panic("unimplemented")
}

// FindAll implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) FindAll() ([]models.Produtos, error) {
	panic("unimplemented")
}

// FindById implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) FindById(id int) (models.Produtos, error) {
	panic("unimplemented")
}

// Update implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) Update(entity models.Produtos) error {
	panic("unimplemented")
}

func NewProdutoRepository(db *configs.Connection) interfaces.RepositoryCrud[models.Produtos] {

	return &ProdutoRepository{

		Repository: *NewRepository(db),
	}
}
