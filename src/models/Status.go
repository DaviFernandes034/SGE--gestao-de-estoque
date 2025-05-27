package models

import "github.com/google/uuid"

type Status struct {
	Id_status uuid.UUID  `json:"status_id"`
	Nome      string      `json:"Nome"`
}
