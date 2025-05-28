package services

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type StatusService struct {
	*BaseService[models.Status]
}

func NewStatusService(rs interfaces.RepositoryCrud[models.Status]) interfaces.ServiceCrud[models.Status]{

	return &StatusService{
		BaseService: NewBaseService[models.Status](rs),
	}

}
