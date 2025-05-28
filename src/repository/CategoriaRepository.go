package repository

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
)

type CategoriasRepository struct {
	Repository
}

// Create implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) Create(entity models.Categorias) error {
	panic("unimplemented")
}

// Delete implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) Delete(entity models.Categorias) error {
	panic("unimplemented")
}

// FindAll implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) FindAll() ([]models.Categorias, error) {
	panic("unimplemented")
}

// FindById implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) FindById(id int) (models.Categorias, error) {
	panic("unimplemented")
}

// Update implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) Update(entity models.Categorias) error {
	panic("unimplemented")
}

func NewCategoriaRepository(db *configs.Connection) interfaces.RepositoryCrud[models.Categorias] {

	return &CategoriasRepository{
		Repository: *NewRepository(db),
	}
}
