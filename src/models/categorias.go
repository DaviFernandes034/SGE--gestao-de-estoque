package models

import "github.com/google/uuid"

type Categorias struct {
	Id_categoria   uuid.UUID `json:"categoria_id"`
	Nome         string `json:"nome"`
}
