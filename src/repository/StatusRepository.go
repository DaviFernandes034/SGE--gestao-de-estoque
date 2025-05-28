package repository

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
)

type StatusRepository struct {
	Repository
}

// Create implements interfaces.RepositoryCrud.
func (s *StatusRepository) Create(entity models.Status) error {
	panic("unimplemented")
}

// Delete implements interfaces.RepositoryCrud.
func (s *StatusRepository) Delete(entity models.Status) error {
	panic("unimplemented")
}

// FindAll implements interfaces.RepositoryCrud.
func (s *StatusRepository) FindAll() ([]models.Status, error) {
	panic("unimplemented")
}

// FindById implements interfaces.RepositoryCrud.
func (s *StatusRepository) FindById(id int) (models.Status, error) {
	panic("unimplemented")
}

// Update implements interfaces.RepositoryCrud.
func (s *StatusRepository) Update(entity models.Status) error {
	panic("unimplemented")
}

func NewStatusRepository(db *configs.Connection) interfaces.RepositoryCrud[models.Status] {

	return &StatusRepository{
		Repository: *NewRepository(db),
	}
}
